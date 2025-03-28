package getobjects

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// GitHub API URL for repository contents
var allPostUrls []string
var extractedValues []string

// GitHubFile represents the JSON structure of a file in the repo
type GitHubFile struct {
	Name        string `json:"name"`
	DownloadURL string `json:"download_url"`
}

func GetCxObjects(swaggerUrl, apiUrl string) []string {
	log.Println("Fetching object list from GitHub API...")

	// Make API request
	resp, err := http.Get(apiUrl)
	if err != nil {
		log.Fatalf("Failed to fetch GitHub API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: status code %d", resp.StatusCode)
	}

	var files []GitHubFile
	if err := json.NewDecoder(resp.Body).Decode(&files); err != nil {
		log.Fatalf("Failed to parse JSON response: %v", err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name, ".md") {
			fmt.Println("\n--- Scanning file:", file.Name, "---")
			findPostEndpoints(file.DownloadURL)
		}
	}

	// Now loop through allPostUrls and fetch the nested values
	log.Println("\n--- Fetching content from all POST API URLs ---")
	fetchAndExtractValues(swaggerUrl)

	log.Println("\n--- Done ---")
	return extractedValues
}

// findPostEndpoints fetches the file content and searches for POST API endpoints
func findPostEndpoints(fileURL string) {
	if fileURL == "" {
		log.Println("Skipping file: No download URL available")
		return
	}

	resp, err := http.Get(fileURL)
	if err != nil {
		log.Printf("Failed to fetch file: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: Unable to fetch file %s (status code: %d)\n", fileURL, resp.StatusCode)
		return
	}

	// Read file content
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read file content: %v\n", err)
		return
	}

	// Search for "POST /api/v2"
	lines := strings.Split(string(body), "\n")
	for _, line := range lines {
		if strings.Contains(line, "POST /api/v2") {
			// Extract the URL from the Markdown link
			a := strings.Split(line, "](")[1]
			a = strings.Split(a, ")")[0]
			allPostUrls = append(allPostUrls, a)
		}
	}
}
// fetchAndExtractValues loops through allPostUrls and extracts nested values
func fetchAndExtractValues(swaggerUrl string) {
    log.Println("Running Fetch on URL")

    resp, err := http.Get(swaggerUrl)
    if err != nil {
        log.Fatalf("Failed to fetch GitHub API: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        log.Fatalf("Error: status code %d", resp.StatusCode)
    }

    var swaggerFile map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&swaggerFile); err != nil {
        log.Fatalf("Failed to parse JSON response: %v", err)
    }

    paths := swaggerFile["paths"].(map[string]interface{})

    // Use a map to track unique values
    uniqueValues := make(map[string]bool)

    for _, url := range allPostUrls {
        if strings.Contains(url, "#post") {
            log.Println("Fetching URL:", url)
            // Format the path to match the Swagger schema
            a := strings.Split(url, "post-")[1]
            a = strings.ReplaceAll(a, "-", "/")
            a = strings.ReplaceAll(a, "//", "/")
            a = "/" + a

            // Wrap Id in {}
            if strings.Contains(a, "Id") {
                temp := strings.Split(a, "/")
                for i, v := range temp {
                    if strings.Contains(v, "Id") {
                        temp[i] = "{" + v + "}"
                    }
                }
                a = strings.Join(temp, "/")
            }

            b, ok := paths[a].(map[string]interface{})
            if !ok {
                log.Println("Path not found in Swagger schema:", a)
                continue
            }

            post := b["post"].(map[string]interface{})
            responses := post["responses"].(map[string]interface{})
            okResponse, ok := responses["200"].(map[string]interface{})
            if !ok {
                log.Println("Response 200 not found in Swagger schema:", a)
                continue
            }
            schema := okResponse["schema"].(map[string]interface{})

            for _, value := range schema {
                v, ok := value.(string)
                if !ok {
                    continue
                }
                if strings.Contains(v, "#/definitions/") {
                    v = strings.Split(v, "/")[2]
                }

                // Add to extractedValues only if it's not already present
                if !uniqueValues[v] {
                    uniqueValues[v] = true
                    extractedValues = append(extractedValues, v)
                }
            }
        }
    }
}