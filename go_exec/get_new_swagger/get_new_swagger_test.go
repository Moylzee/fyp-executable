package getnewswagger

import (
	"log"
	"os"
	"testing"
)

func getSwaggerUrlEnvironmentVariable() string {
	swaggerUrl := os.Getenv("SWAGGER_URL")
	if swaggerUrl == "" {
		log.Fatal("SWAGGER_URL environment variable is not set")
	}
	return swaggerUrl
}

// Mock function to simulate fetching a Swagger file
func TestFetchSwagger(t *testing.T) {
	log.Println("Starting TestFetchSwagger")

	swaggerUrl := getSwaggerUrlEnvironmentVariable()

	result, err := FetchSwagger(swaggerUrl)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
    if result == nil {
        t.Errorf("Expected non-empty result, got nil")
    }

	log.Println("TestFetchSwagger completed successfully")
}