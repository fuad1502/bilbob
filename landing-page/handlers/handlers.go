package handlers

import (
	"embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	// resources represents the embedded file system containing the ./resources directory
	// and all the files and directories contained in it, recursively.
	//
	//go:embed resources
	resources embed.FS

	// webappUrl is the URL of the webapp. This will be written dynamically depending on environment variables
	// when a client requested urls.js.
	webappUrl = os.Getenv("PROTOCOL") + os.Getenv("WEB_SUBDOMAIN") + os.Getenv("HOSTNAME") + ":" + os.Getenv("WEBAPP_PORT")

	// apiEndpoint is the API endpoint URL. This will be written dynamically depending on environment variables
	// when a client requested urls.js.
	apiEndpoint = os.Getenv("PROTOCOL") + os.Getenv("API_SUBDOMAIN") + os.Getenv("HOSTNAME") + ":" + os.Getenv("API_PORT")
)

// LandingPageHandler will serve static files required for interacting with the landing page.
func LandingPageHandler(w http.ResponseWriter, r *http.Request) {
	addCORSHeader(w, r)
	// Get the requested file name.
	fileName, _ := strings.CutPrefix(r.URL.Path, "/"+os.Getenv("LP_PATH"))
	// If no file is requested, then we will serve the index.html file instead
	if fileName == "" {
		fileName = "index.html"
	}
	// If urls.js is requested, write to response writer directly with the
	// API endpoint and webapp URL as JS constants.
	if fileName == "urls.js" {
		urlsJS := fmt.Sprintf("export const api = '%v'; export const webappUrl = '%v'", apiEndpoint, webappUrl)
		w.Header().Add("Content-Type", "text/javascript")
		w.Write([]byte(urlsJS))
		log.Printf("Served %v!\n", fileName)
		return
	}
	// Get file read/seek interface for the requested file.
	// The file might be stored in a different structure than
	// what is presented through the URL.
	if file, err := getReadSeeker(fileName); err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	} else {
		http.ServeContent(w, r, fileName, time.Now(), file)
		log.Printf("Served %v!\n", fileName)
	}
}

func getReadSeeker(fileName string) (io.ReadSeeker, error) {
	path, err := getFilePath(fileName)
	if err != nil {
		return nil, err
	}
	if os.Getenv("LP_MODE") == "release" {
		file, err := resources.Open(path)
		return file.(io.ReadSeeker), err
	} else {
		return os.Open(path)
	}
}

func addCORSHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", webappUrl)
}

func getFilePath(name string) (string, error) {
	ext, err := getExtension(name)
	if err != nil {
		return "", err
	}
	if ext == "jpg" || ext == "png" {
		path := "resources/images/" + name
		return path, nil
	}
	path := "resources/" + ext + "/" + name
	return path, nil
}

func getExtension(filename string) (string, error) {
	re := regexp.MustCompile(`[\w-_]+[.]([\w-_]+)`)
	sm := re.FindStringSubmatch(filename)
	if len(sm) < 2 {
		return "", fmt.Errorf("Invalid file extension")
	}
	return sm[1], nil
}
