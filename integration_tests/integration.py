import subprocess
import json
import unittest
import os

class UmWarsawClientIntegration(unittest.TestCase):
    def test_get_bus_stop(self):
        api_key = os.environ.get('API_KEY')
        output = subprocess.check_output(['api-um-warsaw-client', '--api-key', api_key,
                                          'getBusStop', 'znana'])
        bus_stop = json.loads(output)
        self.assertEqual(bus_stop['BusID'], '5104', "Incorrect bus id")
        self.assertEqual(bus_stop['Name'], 'Znana', "Incorrect bus name")

if __name__ == "__main__":
    unittest.main()