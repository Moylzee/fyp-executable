package getreferenceobjects

import (
	"log"
	"strings"
)

var (   
    AllObjects []string
)

func GetReferenceCxObjects(swagger map[string]interface{}, cxObjects []string) map[string]interface{} {
	// Get All Reference Process
	AllRefs(swagger, cxObjects)

	log.Println("Found All Objects used by CX as Code - Adding Them to JSON")

	definitionMap := make(map[string]interface{})
	for _, object := range AllObjects {
		definition, exists := swagger[object]
		if !exists {
			continue
		}
		definitionMap[object] = definition
	}

	return definitionMap
}

func AllRefs(swagger map[string]interface{}, cxObjects []string) {
	for _, ref := range cxObjects {
		log.Printf("Finding Reference Objects for %s", ref)
		FindAllRefs(swagger[ref])
		if !Contains(AllObjects, ref) {
			AllObjects = append(AllObjects, ref)
		}
	}
}

func FindAllRefs(schema interface{}) {
	if schemaMap, ok := schema.(map[string]interface{}); ok {
		for key, value := range schemaMap {
			if strings.HasSuffix(key, "$ref") {
				value = strings.TrimPrefix(value.(string), "#/definitions/")
				if !Contains(AllObjects, value.(string)) {
					AllObjects = append(AllObjects, value.(string))
				}
			}
			FindAllRefs(value)
		}
	}
}

func Contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
