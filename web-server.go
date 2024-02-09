package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
)

func loadFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func getMimeType(filename string) (string, error) {
	// Use regex to get the file extension
	re := regexp.MustCompile(`[\w-_]+[.]([\w-_]+)`)
	extension := re.FindStringSubmatch(filename)
	if len(extension) < 2 {
		return "", fmt.Errorf("Invalid file extension")
	}
	// Map the file extension to the mime type
	if extension[1] == "html" || extension[1] == "css" {
		return fmt.Sprintf("text/%s", extension[1]), nil
	} else if extension[1] == "js" {
		return "application/javascript", nil
	}
	return "", fmt.Errorf("Invalid file extension")
}

func landingPageHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Path[1:]
	// If the file name is empty, then we will serve the index.html file
	if fileName == "" {
		fileName = "index.html"
	}
	// Validate the file name and get the mime type
	mimeType, err := getMimeType(fileName)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	// Load the file content
	content, err := loadFile(fileName)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	// Set the content type header
	w.Header().Add("Content-Type", mimeType)
	// Write the content to the response writer
	fmt.Fprint(w, content)
}

func main() {
	http.HandleFunc("/", landingPageHandler)
	fmt.Println("Bilbob Web Server is running on port 8080! 🐱")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
