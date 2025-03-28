ADDED_COUNT = 0
MODIFIED_COUNT = 0
REMOVED_COUNT = 0
NEW_OBJECTS = []

# Function to get enum values from a JSON object
def get_enum_values(json_obj, key_prefix):
    enum_values = []
    for key, value in json_obj.items():
        if key.startswith(key_prefix) and ".enum" in key:
            enum_values.append(value)
    return enum_values

def compare_enums(enum_values_anchor, enum_values_swagger):
    # Added Attributes 
    added = list(set(enum_values_swagger) - set(enum_values_anchor))
    # Removed Attributes
    removed = list(set(enum_values_anchor) - set(enum_values_swagger))
    return added, removed

def remove_properties(diff):
    new_diff = {}  # Create a new dictionary to store the modified structure
    
    for key, value in diff.items():
        if isinstance(value, dict):
            # Handle nested dictionaries
            new_value = remove_properties(value)
            
            # Modify the key if necessary
            new_key = key.replace(".properties.", ".") if ".properties." in key else key
            new_diff[new_key] = new_value
        else:
            # Modify the key if necessary
            new_key = key.replace(".properties.", ".") if ".properties." in key else key
            new_diff[new_key] = value
            
    return new_diff

def compare_attributes(anchor_value, swagger_value, attribute_type):
    if anchor_value is None:  # Added Attributes
        return {
            f'added_{attribute_type}': swagger_value
        }
    elif swagger_value is None:  # Removed Attributes
        return {
            f'deleted_{attribute_type}': anchor_value
        }
    elif anchor_value is not None and swagger_value is not None and anchor_value != swagger_value:  # Modified Attributes
        return {
            f'old_{attribute_type}': anchor_value,
            f'new_{attribute_type}': swagger_value,
        }
    return None

# Function to compare two flattened JSON objects with naming convention handling
def compare_jsons(anchor, swagger):
    diff = {}
    global ADDED_COUNT, MODIFIED_COUNT, REMOVED_COUNT, NEW_OBJECTS

    all_keys = set(anchor.keys()).union(set(swagger.keys()))
    attribute_types = ["description", "readOnly", "example", "format"]

    for key in all_keys:
        value1, value2 = anchor.get(key), swagger.get(key)

        # Skip Attributes that have selfUri 
        if ".selfUri." in key:
            continue

        # Handle Enums Appropriately 
        if ".enum" in key:
            base_key = key.split(".enum")[0]
            added, removed = compare_enums(get_enum_values(anchor, base_key), get_enum_values(swagger, base_key))

            # Only include non-empty lists
            enum_diff = {"new_enum": get_enum_values(swagger, base_key)}
            if added:
                enum_diff["added"] = added
            if removed:
                enum_diff["deleted"] = removed
            if added or removed:  # Only add key if there's a change
                diff[base_key] = enum_diff
            continue

        # Dynamic attribute comparison
        for attr_type in attribute_types:
            if key.endswith(f".{attr_type}"):
                result = compare_attributes(value1, value2, attr_type)
                if result:
                    if any(k.startswith("added_") for k in result.keys()):
                        ADDED_COUNT += 1
                    elif any(k.startswith("deleted_") for k in result.keys()):
                        REMOVED_COUNT += 1
                    elif any(k.startswith("old_") for k in result.keys()):
                        MODIFIED_COUNT += 1

                    diff[key] = result
                break

        if key.endswith(".type"):
            # Functionality to Check for completely new objects
            if ".type" in key and len(key.split(".")) == 2:
                if value1 is None: # Added Attributes
                    ADDED_COUNT += 1
                    NEW_OBJECTS.append(key.split(".")[0])
                    diff[key] = {
                        'added_object': key.split(".")[0],
                        'added_type': value2
                    }
                elif value2 is None: # Removed Attributes
                    REMOVED_COUNT += 1
                    diff[key] = {
                        'deleted_object': key.split(".")[0],
                        'deleted_type': value1
                    }
            else: # Functionality to Check for modified Attributes
                if value1 is None: # Added Attributes
                    ADDED_COUNT += 1
                    diff[key] = {
                        'added_type': value2
                    }
                elif value2 is None: # Removed Attributes
                    REMOVED_COUNT += 1
                    diff[key] = {
                        'deleted_type': value1
                    }
                elif value1 is not None and value2 is not None and value1 != value2: # Modified Attributes
                    MODIFIED_COUNT += 1
                    diff[key] = {
                        'old_type': value1,
                        'new_type': value2
                    }
            continue

    diff = remove_properties(diff)
    return diff

def compare_files(anchor, latest_swagger):
    print("Comparing JSON files...")

    differences = compare_jsons(anchor, latest_swagger)

    print("Comparison complete")

    if len(differences) == 0:
        print("No differences found.")
        return None, None, None, None, None
    
    print("Differences found")

    return differences, ADDED_COUNT, MODIFIED_COUNT, REMOVED_COUNT, NEW_OBJECTS
