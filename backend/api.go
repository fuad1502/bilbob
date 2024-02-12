package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	log.SetPrefix("[Bilbob API]: ")

	// Connect to the database
	log.Println("Connecting to the database...")
	safeDB, err := ConnectPGDB(os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to the database!")
	defer safeDB.db.Close()

	// Create a new router
	router := gin.Default()

	// Add CORS middleware
	router.Use(CORSMiddleware())

	// Add users route
	router.GET("/users/:username", checkUser(safeDB))
	router.POST("/users", addUser(safeDB))

	// Run the server
	log.Println("Web service running")
	router.Run(":8080")
}
