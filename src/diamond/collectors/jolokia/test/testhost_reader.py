#!/usr/bin/python
# coding=utf-8
################################################################################

import json
from test import CollectorTestCase
from test import get_collector_config
from test import unittest

import host_reader
from host_reader import StandaloneHostReader, KubernetesHostReader, NerveHostReader
from mock import Mock
from mock import patch


################################################################################

class TestStandaloneHostReader(CollectorTestCase):
    def setUp(self):
        self.host_reader = StandaloneHostReader()

    def test_import(self):
        self.assertTrue(StandaloneHostReader)

    def test_should_return_single_host(self):
        config = {'host': '10.1.1.2'}
        self.host_reader.configure(config)

        expected = {"10.1.1.2": "10.1.1.2"}
        actual = self.host_reader.read()
        self.assertEquals(expected, actual)


class TestNerveHostReader(CollectorTestCase):
    def setUp(self):
        self.config = get_collector_config('JolokiaCollector', {})
        self.host_reader = NerveHostReader()

    def test_import(self):
        self.assertTrue(NerveHostReader)

    def test_should_filter_hosts_success(self):
        def se():
            return json.loads(self.getFixture('nerve.json').getvalue()), None

        with patch("nerve.read", Mock(side_effect=se)):
            self.config['spec'] = {
                'host_id_regex': '^(cassandra_[\\w_-]+)',
            }
            self.host_reader.configure(self.config)
            actual = self.host_reader.read()
            expected = {
                "cassandra_dev.main.norcal-devc:10.93.118.202.9042": "10.40.19.94",
                "cassandra_dev.main.norcal-devc:10.93.118.202.9042.v2.new": "10.93.118.202",
                "cassandra_dev.main.pnw-devc:10.93.118.202.9042": "10.40.19.94",
                "cassandra_dev.main.pnw-devc:10.93.118.202.9042.v2.new": "10.93.118.202"
            }
            self.assertEquals(expected, actual)

    def test_should_not_fail_when_spec_absent(self):
        def se():
            return json.loads(self.getFixture('nerve.json').getvalue()), None

        with patch("nerve.read", Mock(side_effect=se)):
            self.host_reader.configure(self.config)
            actual = self.host_reader.read()
            self.assertEquals({}, actual)

    def test_should_return_empty_list_when_regex_does_not_match(self):
        def se():
            return json.loads(self.getFixture('nerve.json').getvalue()), None

        with patch("nerve.read", Mock(side_effect=se)):
            self.config['spec'] = {
                'host_id_regex': '^(should_not_match_[\\w_-]+)',
            }
            self.host_reader.configure(self.config)
            actual = self.host_reader.read()
            self.assertEquals({}, actual)

    def test_should_not_fail_when_nerve_read_error(self):
        def se():
            return None, IOError("some error")

        with patch("nerve.read", Mock(side_effect=se)):
            self.config['spec'] = {
                'host_id_regex': '^(cassandra_[\\w_-]+)',
            }
            self.host_reader.configure(self.config)
            actual = self.host_reader.read()
            self.assertEquals({}, actual)


class TestKubernetesHostReader(CollectorTestCase):
    def setUp(self):
        self.config = get_collector_config('JolokiaCollector', {})
        self.host_reader = KubernetesHostReader()

    def test_import(self):
        self.assertTrue(KubernetesHostReader)

    def test_should_filter_hosts_success(self):
        def se():
            return json.loads(self.getFixture('pods.json').getvalue()), None

        with patch("kubernetes.Kubelet.list_pods", Mock(side_effect=se)):
            self.config['spec'] = {
                'label_selector': {
                    "yelp.com/paasta_service": "kafka-operator",
                    "yelp.com/paasta_instance": "main"
                }}
            self.host_reader.configure(self.config)
            actual = self.host_reader.read()
            self.assertEquals({"172.23.0.42": "172.23.0.42"}, actual)

    def test_should_filter_hosts_partial_match(self):
        def se():
            return json.loads(self.getFixture('pods.json').getvalue()), None

        with patch("kubernetes.Kubelet.list_pods", Mock(side_effect=se)):
            self.config['spec'] = {
                'label_selector': {
                    "yelp.com/paasta_service": "not_present",
                    "yelp.com/paasta_instance": "main"
                }}
            self.host_reader.configure(self.config)
            actual = self.host_reader.read()
            self.assertEquals({}, actual)

    def test_should_not_fail_when_spec_absent(self):
        def se():
            return json.loads(self.getFixture('pods.json').getvalue()), None

        with patch("kubernetes.Kubelet.list_pods", Mock(side_effect=se)):
            self.host_reader.configure(self.config)
            actual = self.host_reader.read()
            self.assertEquals({}, actual)

    def test_should_return_empty_list_when_no_label_selectors(self):
        def se():
            return json.loads(self.getFixture('pods.json').getvalue()), None

        with patch("kubernetes.Kubelet.list_pods", Mock(side_effect=se)):
            self.config['spec'] = {'label_selector': {}}
            self.host_reader.configure(self.config)
            actual = self.host_reader.read()
            self.assertEquals({}, actual)

    def test_should_not_fail_when_kubelet_is_not_reachable(self):
        def se():
            return None, IOError("some error")

        with patch("kubernetes.Kubelet.list_pods", Mock(side_effect=se)):
            self.config['spec'] = {
                'label_selector': {
                    "yelp.com/paasta_service": "kafka-operator",
                    "yelp.com/paasta_instance": "main"
                }}
            self.host_reader.configure(self.config)
            actual = self.host_reader.read()
            self.assertEquals({}, actual)


class TestHostReader(CollectorTestCase):
    def test_should_get_by_mode(self):
        self.assertEqual(type(KubernetesHostReader()), type(host_reader.get_by_mode('kubernetes')))
        self.assertEqual(type(NerveHostReader()), type(host_reader.get_by_mode('nerve')))
        self.assertEqual(type(StandaloneHostReader()), type(host_reader.get_by_mode('standalone')))

    def test_should_return_default_reader_for_invalid_mode(self):
        self.assertEqual(type(StandaloneHostReader()), type(host_reader.get_by_mode('random')))


################################################################################


if __name__ == "__main__":
    unittest.main()
