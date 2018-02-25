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

    def test_get_lines_at_bus_stop(self):
        api_key = os.environ.get('API_KEY')
        output = subprocess.check_output(['api-um-warsaw-client', '--api-key', api_key,
                                          'getLinesAtBusStop', '5104', '01'])
        lines = json.loads(output)
        self.assertEqual(len(lines), 3)
        self.assertSequenceEqual(lines, ['129', '155', '167'])

    def test_get_time_table(self):
        api_key = os.environ.get('API_KEY')
        output = subprocess.check_output(['api-um-warsaw-client', '--api-key', api_key,
                                          'getTimeTable', '5104', '01', '155'])
        timeTable = json.loads(output)
        self.assertGreater(len(timeTable), 0)
        for record in timeTable:
            self.assertIn("Brigade", record)
            self.assertIn("Direction", record)
            self.assertIn("Time", record)
            self.assertIsNotNone(record['Brigade'])
            self.assertIsNotNone(record['Direction'])
            self.assertIsNotNone(record['Time'])

if __name__ == "__main__":
    unittest.main()