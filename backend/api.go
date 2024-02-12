package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"sync"
)

// SafeDB is a wrapper around sql.DB that provides a mutex to make it safe for concurrent use
type SafeDB struct {
	mu sync.Mutex
	db *sql.DB
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func checkLogin(safeDB *SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"result": "success"})
	}
}

func main() {
	log.SetPrefix("[Bilbob API]: ")

	// Connect to the "postgres" database
	log.Println("Connecting to the database...")
	connStr := "host=data user=postgres password=secret dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the connection is working
	log.Println("Pinging the database...")
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	log.Println("Connected to the database")

	// Create a new SafeDB
	safeDB := &SafeDB{db: db}

	// Create a new router
	router := gin.Default()
	// Use the CORS middleware
	router.Use(CORSMiddleware())
	// Add the login route
	router.GET("/login", checkLogin(safeDB))
	// Run the server
	router.Run(":8080")
}
