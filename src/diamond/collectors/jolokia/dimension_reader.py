import abc
import re
import string

import kubernetes
import nerve


class DimensionReader(object):
    __metaclass__ = abc.ABCMeta

    @abc.abstractmethod
    def name(self):
        """
        Return the name of the dimension reader
        :return: name of the dimension reader
        """
        pass

    @abc.abstractmethod
    def configure(self, conf):
        """
        Configures the dimension reader
        :param conf:
        :return:
        """
        pass

    @abc.abstractmethod
    def read(self, hosts):
        """
        Computes the dimension for the given hosts
        :param hosts: dict of host identifiers to host address
        :return: a dict of host identifier to their dimensions
        """
        pass


class NoopDimensionReader(DimensionReader):
    """
    A dimension reader which always returns an empty dictionary
    """

    def name(self):
        return ""

    def configure(self, conf):
        pass

    def read(self, hosts):
        return {}


class LabelRegex(object):
    def __init__(self, label, regex):
        self.label = label
        self.regex = regex


class KubernetesDimensionReader(DimensionReader):
    """
    A dimension reader which queries the ```/pods``` and creates dimension from the label
    annotations on the pods. This readers need the final dimension_name, label_name and a
    regex to convert label annotation to dimensions.

    Example:
    ```
        "spec": {
            "dimensions": {
                "kubernetes": {
                    "paasta_service": {
                        "paasta.yelp.com/service": "[a-z]*"
                    },
                    "paasta_instance": {
                        "paasta.yelp.com/instance": ".*"
                    }
                }
            }
        }
    ```
    In above config, the kubernetes dimension reader is configured to create 2 dimensions.
    The first dimension name is ```paasta_service``` and dimension value is created by
    apply regex ```[a-z]*``` on the value of label ```paasta.yelp.com/service``` in the
    pods metadata.
    So dimension extraction can be expressed in the following format
    ```
        "kubernetes" : {
            "${dimension_name}" : {
                "${label_name}" : "${regex}"
            }
        }
    ```
    """

    def __init__(self):
        self.dim_generators = {}
        self.kubelet = kubernetes.Kubelet()

    def name(self):
        return "kubernetes"

    def configure(self, conf):
        self.dim_generators = self.create_dim_generators(conf)

    def read(self, hosts):
        response, err = self.kubelet.list_pods()
        host_dimension = {}
        if err is not None or 'items' not in response:
            return host_dimension
        pods = response.get('items', [])
        for pod in pods:
            pod_ip = pod.get('status', {}).get('podIP')
            # find the host identifier with the given address
            host_id = _find_host_id_by_ip(hosts, pod_ip)
            if pod_ip is None or host_id is None:
                continue
            labels = pod.get('metadata', {}).get('labels', {})
            host_dimension[host_id] = self.generate_dimension(labels, self.dim_generators)
        return host_dimension

    def create_dim_generators(self, dimension_regex):
        dim_compile_rx = {}
        for dim, generator in dimension_regex.items():
            for label, regex in generator.items():
                dim_compile_rx[dim] = LabelRegex(label=label, regex=re.compile(regex))
        return dim_compile_rx

    def generate_dimension(self, pod_labels, dim_generators):
        generated = {}
        for dim, label_regex in dim_generators.items():
            label_value = pod_labels.get(label_regex.label)
            if label_value is not None:
                matches = label_regex.regex.findall(label_value)
                if len(matches) > 0:
                    generated[str(dim)] = string.replace(matches[0], "--", "_", -1)
        return generated


class NerveDimensionReader(DimensionReader):
    """
   A dimension reader which readers the nerve config file and creates dimension from the service
   identifiers. This readers need the dimension_name and a regex to extract dimension from
   the service identifier

   Example:
   ```
       {
            'spec': {
                'dimensions': {
                    'nerve': {
                        'cassandra_cluster': '^cassandra_([\\w_-]+).',
                        'cassandra_instance': '^cassandra_[\\w_-]+.([\\w_-]+).',
                    }
                },
            }
        })
   ```
   In above config, the nerve dimension reader is configured to create 2 dimensions.
   The first dimension name is ```cassandra_cluster``` and dimension value is created by
   apply regex ```^cassandra_([\\w_-]+).'``` on the value of label ```cassandra_cluster``` in the
   pods metadata.
   So dimension extraction can be expressed in the following format
   ```
       "nerve" : {
           "${dimension_name}" : "${regex}"
       }
   ```
   Note the regex can also contains groups. For example, for regex ```^cassandra_([\\w_-]+).``` and
   service identifier ```cassandra_dev.main.norcal-devc:10.93.118.202.9042.v2.new```, the value
   ```dev``` will be extracted as the dimension value
   """

    def __init__(self):
        super(NerveDimensionReader, self).__init__()
        self.dim_generators = {}

    def pr(self):
        return self.dim_generators

    def name(self):
        return "nerve"

    def configure(self, conf):
        self.dim_generators = self.create_dim_generators(conf)

    def read(self, hosts):
        values, err = nerve.read()
        if err is not None:
            return {}
        services = values.get('services', {})
        host_dimensions = {}
        for host_identifier, host_ip in hosts.items():
            value = services.get(host_identifier)
            if value is None:
                continue
            host_dimensions[host_identifier] = self.generate_dimensions(host_identifier, self.dim_generators)
        return host_dimensions

    def create_dim_generators(self, dimension_regex):
        dim_compile_rx = {}
        for dim, regex in dimension_regex.items():
            dim_compile_rx[dim] = re.compile(regex)
        return dim_compile_rx

    def generate_dimensions(self, host_id, dim_generators):
        generated = {}
        for dim, regex in dim_generators.items():
            matches = regex.findall(host_id)
            if len(matches) > 0:
                generated[dim] = matches[0]
        return generated


class CompositeDimensionReader(DimensionReader):
    """
    A dimension reader that readers dimensions from multiple readers and
    merges them into a single host-dimension dictionary. This readers does
    nothing of no readers are configured
    """

    def __init__(self):
        super(CompositeDimensionReader, self).__init__()
        self.readers = []

    def name(self):
        return "composite"

    def configure(self, conf):
        readers = []
        for name, reader_conf in conf.items():
            reader = get_reader(name)
            if reader is None:
                continue
            reader.configure(reader_conf)
            readers.append(reader)
        self.readers = readers

    def read(self, hosts):
        dims = {}
        for reader in self.readers:
            self.merge(dims, reader.read(hosts))
        return dims

    @staticmethod
    def merge(dim1, dim2):
        for host, dims in dim2.items():
            dim1.setdefault(host, {}).update(dims)


REGISTRY = {
    'kubernetes': KubernetesDimensionReader(),
    'nerve': NerveDimensionReader(),
}

DEFAULT_DIMENSION_READER = NoopDimensionReader()


def get_reader(name):
    """
    Returns the dimension reader for the given name
    :param name: name of the dimension readereidfcciknhlfduigieigndrknvcrvjckvnjjhbiedfhf

    :return: dimension reader by name or ```DEFAULT_DIMENSION_READER```
    """
    return REGISTRY.get(name, DEFAULT_DIMENSION_READER)


def _find_host_id_by_ip(hosts, ip):
    for host_id, host_ip in hosts.items():
        if ip == host_ip:
            return host_id
    return None