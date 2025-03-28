package prepareJson

import (
	"log"

	"github.com/jeremywohl/flatten"
)

func FinalPrepSwagger(swagger map[string]interface{}) (map[string]interface{}, error) {
	log.Println("Flattening Latest Swagger")

	finalSwagger, err := flatten.Flatten(swagger, "", flatten.DotStyle)
	if err != nil {
		return nil, err
	}

	log.Println("Flattened Latest Swagger")
	return finalSwagger, nil
}