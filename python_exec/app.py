import json
import os
from datetime import datetime
from utils.common import ReadFromS3, unflatten_dict, upload_file_to_s3
from compare_swaggers_python.compare import compare_files
from post_to_chat.post import post_to_chat

def write_to_file(file_path, content):
    with open(file_path, 'w') as file:
        file.write(content)
    

def main():
    anchor_swagger_key = '../bucket/anchor_swagger.json'
    latest_swagger_key = '../bucket/final_swagger.json'
    with open(anchor_swagger_key, 'r') as file:
        anchor_swagger = json.load(file)
    with open(latest_swagger_key, 'r') as file:
        latest_swagger = json.load(file)


    result, added_count, modified_count, removed_count, new_objects = compare_files(anchor_swagger, latest_swagger)
    
    if result is None:
        NO_CHANGES_MESSAGE = f'-------------------------------\nNo Changes Detected In Swagger' 
        print(NO_CHANGES_MESSAGE)
        empty_json = {}
        write_to_file('../bucket/summary.json', json.dumps(empty_json, indent=4))
        return 

    total = added_count + modified_count + removed_count

    SUMMARY_MESSAGE = f'----------------------------------\nChanges Detected In Swagger\n' \
    f'Number of Changes Detected: {total}\n' \
    f'Added: {added_count}\n' \
    f'Modified: {modified_count}\n' \
    f'Removed: {removed_count}\n' \
    

    if new_objects:
        SUMMARY_MESSAGE += f'New Objects: [{", ".join(new_objects)}]\n'

    print(SUMMARY_MESSAGE)

    write_to_file('../bucket/summary.json', json.dumps(result, indent=4))

    write_to_file('../bucket/anchor_swagger.json', json.dumps(latest_swagger, indent=4))

    return

main()