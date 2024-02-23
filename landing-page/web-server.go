package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
)

var webappUrl = os.Getenv("PROTOCOL") + os.Getenv("HOSTNAME") + ":" + os.Getenv("WEBAPP_PORT")
var apiEndpoint = os.Getenv("PROTOCOL") + os.Getenv("HOSTNAME") + ":" + os.Getenv("API_PORT")

func loadFile(filename string) (string, string, error) {
	ext, err := getExtension(filename)
	if err != nil {
		return "", "", err
	}
	mimeType, err := getMimeType(ext)
	if err != nil {
		return "", "", err
	}
	filePath := "resources/" + ext + "/" + filename
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", "", err
	}
	return string(content), mimeType, nil
}

func getExtension(filename string) (string, error) {
	re := regexp.MustCompile(`[\w-_]+[.]([\w-_]+)`)
	sm := re.FindStringSubmatch(filename)
	if len(sm) < 2 {
		return "", fmt.Errorf("Invalid file extension")
	}
	return sm[1], nil
}

func getMimeType(ext string) (string, error) {
	if ext == "html" || ext == "css" {
		return fmt.Sprintf("text/%s", ext), nil
	} else if ext == "js" {
		return "application/javascript", nil
	}
	return "", fmt.Errorf("Unsupported file extension")
}

func landingPageHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Path[1:]
	// If the file name is empty, then we will serve the index.html file
	if fileName == "" {
		fileName = "index.html"
	}
	// Load the file and get the mime type
	content, mimeType, err := loadFile(fileName)
	if err != nil {
		http.Error(w, "", http.StatusNotFound)
		log.Printf("Unable to serve %v: %v", fileName, err)
		return
	}
	// Set the content type header
	w.Header().Add("Content-Type", mimeType)
	// Write the content to the response writer
	fmt.Fprint(w, content)
	log.Printf("Served %v!\n", fileName)
}

func wrapCORSHandler(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", webappUrl)
		handler(w, r)
	}
}

func overwriteApiEndpointJS() error {
	content := fmt.Sprintf("export var api = '%v';", apiEndpoint)
	return os.WriteFile("resources/js/api-endpoint.js", []byte(content), 0644)
}

func main() {
	log.SetPrefix("[Bilbob WebServer]: ")
	if err := overwriteApiEndpointJS(); err != nil {
		panic(err)
	}
	http.HandleFunc("/", wrapCORSHandler(landingPageHandler))
	log.Println("Bilbob Web Server is running on port 8080! 🐱")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
