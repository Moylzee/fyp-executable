package main

import (
	"encoding/json"
	getnewswagger "go_exec/get_new_swagger"
	getobjects "go_exec/get_objects"
	getreferenceobjects "go_exec/get_reference_objects"
	prepareJson "go_exec/prepare_json"
	"log"
	"os"
)


const (
	swaggerUrl = "https://s3.dualstack.us-east-1.amazonaws.com/inin-prod-api/us-east-1/public-api-v2/swagger-schema/publicapi-v2-latest.json"
	apiUrl = "https://api.github.com/repos/MyPureCloud/terraform-provider-genesyscloud/contents/docs/resources"
)

func main() {
	log.Println("Starting Lambda Function")

	log.Println("Fetching Swagger File")

	rawSwagger, err := getnewswagger.FetchSwagger(swaggerUrl)
	if err != nil {
		log.Fatalf("Failed to fetch swagger: %v", err)
	}

	log.Println("Swagger File Fetched")

	log.Println("Fetching Objects")
	cxObjects := getobjects.GetCxObjects(swaggerUrl, apiUrl)

	log.Println("Objects Fetched")

	log.Println("Preparing Final Swagger")
	preprocessedSwagger := getreferenceobjects.GetReferenceCxObjects(rawSwagger, cxObjects)

	finalSwagger, err := prepareJson.FinalPrepSwagger(preprocessedSwagger)
	if err != nil {
		log.Fatalf("Failed to prepare final swagger: %v", err)
	}

	log.Println("Final Swagger Prepared")

	log.Println("Uploading Final Swagger to S3")
	data, err := json.Marshal(finalSwagger)
	if err != nil {
		log.Fatalf("Failed to marshal final swagger: %v", err)
	}

	filePath := "../bucket/final_swagger.json"
	err = WriteToFile(filePath, data)
	if err != nil {
		log.Fatalf("Failed to write final swagger to file: %v", err)
	}

	log.Println("Final Swagger Uploaded to S3")
}

func WriteToFile(filePath string, data []byte) error {
	// Create a new file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the data to the file
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

