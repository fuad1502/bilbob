package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fuad1502/bilbob/landing-page/handlers"
)

func main() {
	log.SetPrefix("[Bilbob WebServer]: ")
	http.HandleFunc("/"+os.Getenv("LP_PATH"), handlers.LandingPageHandler)
	log.Println("Bilbob Web Server is running on port 8080! 🐱")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
