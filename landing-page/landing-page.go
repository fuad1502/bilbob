package main

import (
	"github.com/fuad1502/bilbob/landing-page/handlers"
	"log"
	"net/http"
)

func main() {
	log.SetPrefix("[Bilbob WebServer]: ")
	http.HandleFunc("/", handlers.LandingPageHandler)
	log.Println("Bilbob Web Server is running on port 8080! 🐱")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
