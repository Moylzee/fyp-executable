import unittest
import logging
from python_lambda.compare_swaggers_python.compare import (
    remove_properties,
    compare_attributes,
    compare_jsons,
)

# Configure logging
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(message)s')

class TestCompareFunctions(unittest.TestCase):
    def setUp(self):
        logging.info(f"Starting test: {self._testMethodName}")

    def tearDown(self):
        logging.info(f"Finished test: {self._testMethodName}")

    def test_remove_properties(self):
        diff = {
            "object1.properties.name": "value1",
            "object2.properties.age": "value2"
        }
        result = remove_properties(diff)
        expected = {
            "object1.name": "value1",
            "object2.age": "value2"
        }
        self.assertEqual(result, expected)

    def test_compare_attributes(self):
        result = compare_attributes(None, "new_value", "description")
        self.assertEqual(result, {"added_description": "new_value"})

        result = compare_attributes("old_value", None, "description")
        self.assertEqual(result, {"deleted_description": "old_value"})

        result = compare_attributes("old_value", "new_value", "description")
        self.assertEqual(result, {
            "old_description": "old_value",
            "new_description": "new_value"
        })

        result = compare_attributes("same_value", "same_value", "description")
        self.assertIsNone(result)

    def test_compare_jsons(self):
        anchor = {
            "object1.description": "old_description",
        }
        swagger = {
            "object1.description": "new_description",
        }

        result = compare_jsons(anchor, swagger)

        expected = {
            "object1.description": {
                "old_description": "old_description",
                "new_description": "new_description"
            }
        }

        self.assertEqual(result, expected)

if __name__ == "__main__":
    unittest.main()
