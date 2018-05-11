import unittest
from main import cycle_length

class TestMain(unittest.TestCase):

    def test_cycle_length(self):
        self.assertEqual(cycle_length(22), 16)

if __name__ == '__main__':
    unittest.main()
