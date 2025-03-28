package getreferenceobjects

import (
	"log"
	"testing"
)

func TestGetRefObjects(t *testing.T) {
	log.Println("Starting TestGetRefObjects")

	result := GetReferenceCxObjects(testSwagger, []string{"Action", "CallableTimeSet"})
	if len(result) == 0 {
		t.Errorf("Expected non-empty result, got nil or empty map")
	}

	if _, exists := result["Action"]; !exists {
		t.Errorf("Expected 'Action' to be in the result map")
	}
	if _, exists := result["CallableTimeSet"]; !exists {
		t.Errorf("Expected 'CallableTimeSet' to be in the result map")
	}
	if _, exists := result["CallableTime"]; !exists {
		t.Errorf("Expected 'CallableTime' to be in the result map")
	}

	log.Println("TestGetRefObjects completed successfully")
}

func TestContainsFunction(t *testing.T) {
	log.Println("Starting TestContainsFunction")

	slice := []string{"Action", "CallableTimeSet"}
	if !Contains(slice, "Action") {
		t.Errorf("Expected 'Action' to be found in the slice")
	}
	if Contains(slice, "NonExistent") {
		t.Errorf("Expected 'NonExistent' to not be found in the slice")
	}

	log.Println("TestContainsFunction completed successfully")
}

var testSwagger = map[string]interface{}{
	"Action": map[string]interface{}{
		"properties": map[string]interface{}{
			"category": map[string]interface{}{
				"description": "Category of Action",
				"type": "string",
			},
			"config": map[string]interface{}{
				"$ref": "#/definitions/ActionConfig",
				"description": "Configuration to support request and response processing",
			},
			"contract": map[string]interface{}{
				"$ref": "#/definitions/ActionContract",
				"description": "Action contract",
			},
			"id": map[string]interface{}{
				"description": "The globally unique identifier for the object.",
				"readOnly": true,
				"type": "string",
			},
			"integrationId": map[string]interface{}{
				"description": "The ID of the integration for which this action is associated",
				"type": "string",
			},
			"name": map[string]interface{}{
				"type": "string",
			},
			"secure": map[string]interface{}{
				"description": "Indication of whether or not the action is designed to accept sensitive data",
				"type": "boolean",
			},
			"selfUri": map[string]interface{}{
				"description": "The URI for this object",
				"format": "uri",
				"readOnly": true,
				"type": "string",
			},
			"version": map[string]interface{}{
				"description": "Version of this action",
				"format": "int32",
				"type": "integer",
			},
		},
		"type": "object",
	},
	"CallableTime": map[string]interface{}{
		"properties": map[string]interface{}{
			"timeSlots": map[string]interface{}{
				"description": "The time intervals for which it is acceptable to place outbound calls.",
				"items": map[string]interface{}{
					"$ref": "#/definitions/CampaignTimeSlot",
				},
				"type": "array",
			},
			"timeZoneId": map[string]interface{}{
				"description": "The time zone for the time slots; for example, Africa/Abidjan",
				"example":     "Africa/Abidjan",
				"type":        "string",
			},
		},
		"required": []string{"timeSlots", "timeZoneId"},
		"type":     "object",
	},
	"CallableTimeSet": map[string]interface{}{
		"properties": map[string]interface{}{
			"callableTimes": map[string]interface{}{
				"description": "The list of CallableTimes for which it is acceptable to place outbound calls.",
				"items": map[string]interface{}{
					"$ref": "#/definitions/CallableTime",
				},
				"type": "array",
			},
			"dateCreated": map[string]interface{}{
				"description": "Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z",
				"format":      "date-time",
				"readOnly":    true,
				"type":        "string",
			},
			"dateModified": map[string]interface{}{
				"description": "Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z",
				"format":      "date-time",
				"readOnly":    true,
				"type":        "string",
			},
			"id": map[string]interface{}{
				"description": "The globally unique identifier for the object.",
				"readOnly":    true,
				"type":        "string",
			},
			"name": map[string]interface{}{
				"description": "The name of the CallableTimeSet.",
				"type":        "string",
			},
			"selfUri": map[string]interface{}{
				"description": "The URI for this object",
				"format":      "uri",
				"readOnly":    true,
				"type":        "string",
			},
			"version": map[string]interface{}{
				"description": "Required for updates, must match the version number of the most recent update",
				"format":      "int32",
				"type":        "integer",
			},
		},
		"required": []string{"callableTimes", "name"},
		"type":     "object",
	},
}

