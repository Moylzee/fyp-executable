package prepareJson

import (
    "log"
    "reflect"
    "testing"
)

func TestPrepareJSON(t *testing.T) {
    log.Println("Starting TestPrepareJSON")

    swagger := map[string]interface{}{
        "Action": map[string]interface{}{
            "properties": map[string]interface{}{
                "category": map[string]interface{}{
                    "description": "Category of Action",
                    "type":        "string",
                },
                "config": map[string]interface{}{
                    "$ref":        "#/definitions/ActionConfig",
                    "description": "Configuration to support request and response processing",
                },
            },
        },
    }

    // Call the function being tested
    result, err := FinalPrepSwagger(swagger)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
    if result == nil {
        t.Errorf("Expected non-empty result, got nil")
    }

    // Define the expected flattened Swagger
    expected := map[string]interface{}{
        "Action.properties.category.description": "Category of Action",
        "Action.properties.category.type":        "string",
        "Action.properties.config.$ref":          "#/definitions/ActionConfig",
        "Action.properties.config.description":   "Configuration to support request and response processing",
    }

    // Compare the result with the expected map
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Expected flattened Swagger: %v, got: %v", expected, result)
    }

    log.Println("TestPrepareJSON completed successfully")
}