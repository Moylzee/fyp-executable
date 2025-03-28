import unittest
from python_lambda.utils.common import unflatten_dict, ReadFromS3
from python_lambda import getEnvironmentVariables
import logging
import sys

logging.basicConfig(level=logging.INFO, stream=sys.stdout)

class TestReadFromS3(unittest.TestCase):
    def setUp(self):
        logging.info("Setting up test environment")
        self.environment_variables = getEnvironmentVariables()
        logging.info(f"Environment variables: {self.environment_variables}")

    def test_read_from_s3(self):
        logging.info("Running test_read_from_s3")

        # Assuming the S3 bucket and object key are set correctly
        anchor_swagger_key = 'anchor_swagger/anchor_swagger.json'

        try:
            data = ReadFromS3(anchor_swagger_key, self.environment_variables['BUCKET_NAME'])
            self.assertIsInstance(data, dict)  # Check if the data is a dictionary
            logging.info("test_read_from_s3 passed")
        except Exception as e:
            logging.error(f"Error in test_read_from_s3: {str(e)}")
            self.fail(f"ReadFromS3 raised an exception: {str(e)}")

        logging.info("test_read_from_s3 completed")

class TestUnflattenDict(unittest.TestCase):
    def test_single_key_value(self):
        logging.info("Running test_single_key_value")

        flat_dict = {"key": "value"}
        expected = {"key": "value"}
        self.assertEqual(unflatten_dict(flat_dict), expected)

        logging.info("test_single_key_value passed")

    def test_nested_keys(self):
        logging.info("Running test_nested_keys")

        flat_dict = {"a.b.c": 1, "a.b.d": 2, "x.y": 3}
        expected = {
            "a": {
                "b": {
                    "c": 1,
                    "d": 2
                }
            },
            "x": {
                "y": 3
            }
        }
        self.assertEqual(unflatten_dict(flat_dict), expected)

        logging.info("test_nested_keys passed")

    def test_custom_separator(self):
        logging.info("Running test_custom_separator")

        flat_dict = {"a-b-c": 1, "x-y": 2}
        expected = {
            "a": {
                "b": {
                    "c": 1
                }
            },
            "x": {
                "y": 2
            }
        }
        self.assertEqual(unflatten_dict(flat_dict, separator='-'), expected)

        logging.info("test_custom_separator passed")

    def test_empty_dict(self):
        logging.info("Running test_empty_dict")

        flat_dict = {}
        expected = {}
        self.assertEqual(unflatten_dict(flat_dict), expected)

        logging.info("test_empty_dict passed")

    def test_no_nesting(self):
        logging.info("Running test_no_nesting")

        flat_dict = {"a": 1, "b": 2}
        expected = {"a": 1, "b": 2}
        self.assertEqual(unflatten_dict(flat_dict), expected)

        logging.info("test_no_nesting passed")

if __name__ == "__main__":
    unittest.main()
