import abc
import re

import kubernetes
import nerve


class HostReader(object):
    __metaclass__ = abc.ABCMeta

    @abc.abstractmethod
    def name(self):
        """
        Returns the name of the host reader
        :return: the name of the host reader
        """
        pass

    @abc.abstractmethod
    def configure(self, conf):
        """
        Configures the host read to read configuration
        :param conf:
        :return:
        """
        pass

    @abc.abstractmethod
    def read(self):
        """
        Returns to a dict of host identifier and host addree
        :return:
        """
        pass


class NerveHostReader(HostReader):
    """
    A host reader that reader form nerve config file.
    The host reader uses regex configured in ```host_id_regex```  field of ```spec```
    to match host identifier in the nerve conf file to read matching hosts
    """

    def __init__(self):
        super(NerveHostReader, self).__init__()
        self.regex = None

    def name(self):
        return "nerve"

    def configure(self, conf):
        rx = conf.get('spec', {}).get('host_id_regex')
        if rx is not None:
            self.regex = re.compile(rx)

    def read(self):
        if self.regex is None:
            return {}

        data, err = nerve.read()
        if err is not None:
            return {}

        hosts = {}
        for name, values in data.get('services', {}).items():
            if self.regex.match(name):
                hosts[name] = values.get('host')
        return hosts


class KubernetesHostReader(HostReader):
    """
    A host reader with query kubelet's ```/pods``` endpoint and filters out hosts matching
    the given ```label_selector```. Label filters is similar to kubernetes
    """

    def __init__(self):
        self.label_selectors = {}
        self.kubelet = kubernetes.Kubelet()

    def name(self):
        return "kubernetes"

    def configure(self, conf):
        self.label_selectors = conf.get('spec', {}).get('label_selector', {})

    def read(self):
        if len(self.label_selectors) == 0:
            return {}
        response, err = self.kubelet.list_pods()
        hosts = {}
        if err is not None:
            return hosts
        pods = response.get('items', [])
        for pod in pods:
            if self._pod_matches_selectors(pod):
                pod_ip = pod.get('status', {}).get('podIP')
                if pod_ip is not None:
                    hosts[pod_ip] = pod_ip
        return hosts

    def _pod_matches_selectors(self, pod):
        """
        check if a pod matches the label_selectors

        :param pod: the given
        :return: true if the given pod has all the labels
        """
        pod_labels = pod.get('metadata', {}).get('labels', {})
        if len(pod_labels) == 0:
            return False
        for name, value in self.label_selectors.items():
            pod_label_value = pod_labels.get(name)
            if pod_label_value is None or pod_label_value != value:
                return False
        return True


class StandaloneHostReader(HostReader):
    """
    A host reader that read the host from the config and simply returns it
    """

    def __init__(self):
        self.hosts = {}

    def name(self):
        return "standalone"

    def configure(self, conf):
        host = conf['host']
        self.hosts = {host: host}

    def read(self):
        return self.hosts


REGISTRY = {
    'kubernetes': KubernetesHostReader(),
    'standalone': StandaloneHostReader(),
    'nerve': NerveHostReader(),
}

DEFAULT_READER = 'standalone'


def get_by_mode(mode):
    """
    Return a host reader registered by the given name
    :param mode:
    :return: host reader by name or default if not is found
    """
    reader = REGISTRY.get(mode)
    if reader is None:
        reader = REGISTRY[DEFAULT_READER]
    return reader
