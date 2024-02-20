package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/fuad1502/bilbob-backend/dbs"
	"github.com/fuad1502/bilbob-backend/middlewares"
	"github.com/fuad1502/bilbob-backend/routes"
)

func main() {
	log.SetPrefix("[Bilbob API]: ")

	// Connect to the database
	log.Println("Connecting to the database...")
	var safeDB *dbs.SafeDB
	var err error
	connected := false
	for retry_count := 0; retry_count < 10; retry_count += 1 {
		safeDB, err = dbs.ConnectPGDB(os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
		if err == nil {
			connected = true
			break
		}
		time.Sleep(time.Second)
		log.Println("Retrying database connection...")
	}
	if !connected {
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
	router.GET("/users/:username", routes.CreateGetUserInfoHandler(safeDB))
	router.POST("/users", routes.CreatePostUserHandler(safeDB))

	// Add follows route
	router.GET("/users/follows", routes.CreateGetFollowsHandler(safeDB))

	// Add posts route
	router.GET("/posts", routes.CreateGetPostsHandler(safeDB))
	router.POST("/posts", routes.CreatePostPostHandler(safeDB))

	// Add authorization route
	router.GET("/authorize", routes.CreateAuthorizeHandler(safeDB))

	// Run the server
	log.Println("Web service running")
	router.Run(":8080")
}
