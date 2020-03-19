import json

NERVE_CONFIG_FILEPATH = '/etc/nerve/nerve.conf.json'


def read():
    try:
        with open(NERVE_CONFIG_FILEPATH, 'r') as f:
            return json.load(f), None
    except Exception as e:
        return None, e
