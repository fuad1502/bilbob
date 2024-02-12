package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	log.SetPrefix("[Bilbob API]: ")

	// Connect to the database
	safeDB, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer safeDB.db.Close()

	// Create a new router
	router := gin.Default()

	// Add CORS middleware
	router.Use(CORSMiddleware())

	// Add users route
	router.GET("/users/:username", checkUser(safeDB))
	router.POST("/users", addUser(safeDB))

	// Run the server
	router.Run(":8080")
}
