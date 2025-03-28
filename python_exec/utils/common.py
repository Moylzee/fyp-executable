import boto3
import json

def ReadFromS3(object_key, bucket_name):
    # Initialize the S3 client
    s3_client = boto3.client('s3')

    try:
        # Fetch the object from S3
        response = s3_client.get_object(Bucket=bucket_name, Key=object_key)
        
        # The 'Body' contains the content of the file
        file_content = response['Body'].read().decode('utf-8')  # Read and decode as a string
        
        # If the file content is JSON, you can load it into a dictionary
        data = json.loads(file_content)
        return data
    
    except Exception as e:
        print(f"Error reading from S3: {str(e)}")
        raise e
    

def unflatten_dict(flat_dict, separator='.'):
    """Unflattens a dictionary with keys separated by a specific separator."""
    unflattened = {}
    for key, value in flat_dict.items():
        parts = key.split(separator)
        d = unflattened
        for part in parts[:-1]:
            d = d.setdefault(part, {})
        d[parts[-1]] = value
    return unflattened

def upload_file_to_s3(file_name: str, content: bytes, bucket: str) -> None:
    """Uploads a file to an S3 bucket."""
    s3_client = boto3.client("s3")
    
    try:
        s3_client.put_object(Bucket=bucket, Key=file_name, Body=content)
    except Exception as e:
        raise e
