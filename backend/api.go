package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"os"

	"github.com/fuad1502/bilbob-backend/dbs"
	"github.com/fuad1502/bilbob-backend/middlewares"
	"github.com/fuad1502/bilbob-backend/routes"
)

func main() {
	log.SetPrefix("[Bilbob API]: ")

	// Connect to the database
	log.Println("Connecting to the database...")
	safeDB, err := dbs.ConnectPGDB(os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to the database!")
	defer safeDB.DB.Close()

	// Create a new router
	router := gin.Default()

	// Add CORS middleware
	router.Use(middlewares.CORSMiddleware())

	// Add error middleware
	router.Use(middlewares.ErrorMiddleware())

	// Add users route
	router.GET("/users/:username/:action", routes.CreateUserActionHandler(safeDB))
	router.POST("/users", routes.CreatePostUserHandler(safeDB))

	// Run the server
	log.Println("Web service running")
	router.Run(":8080")
}
