package getnewswagger

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func FetchSwagger(swaggerUrl string) (map[string]interface{}, error) {
	var data map[string]interface{}

	log.Println("Retrieving Swagger File From URL")
	resp, err := http.Get(swaggerUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, err
	}

	log.Println("Retrieved Swagger File")	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	} 
	
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	// Filter By CX as Code Definitions
	defJson := data["definitions"].(map[string]interface{})

	return defJson, nil
}
