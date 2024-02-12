package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type UserSignup struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Animal   string `json:"animal" binding:"required"`
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

// checkUser is a handler that checks if a user exists in the database
func checkUser(safeDB *SafeDB) gin.HandlerFunc {
	// Prepare SQL statemtn for checking if a user exists
	safeDB.mu.Lock()
	defer safeDB.mu.Unlock()
	stmt, err := safeDB.db.Prepare("SELECT username FROM Users WHERE username = $1")
	if err != nil {
		log.Fatal(err)
	}

	return func(c *gin.Context) {
		safeDB.mu.Lock()
		defer safeDB.mu.Unlock()

		// Get the username from the URL
		username := c.Param("username")

		// Query the database for the user
		row := stmt.QueryRow(username)
		if err := row.Scan(&username); err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusOK, gin.H{"result": "user does not exist"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": "user exists"})
	}
}

// addUser is a handler that adds a user to the database
func addUser(safeDB *SafeDB) gin.HandlerFunc {
	// Prepare SQL statement for adding a user
	safeDB.mu.Lock()
	defer safeDB.mu.Unlock()
	stmt, err := safeDB.db.Prepare("INSERT INTO Users (username, password, name, animal) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Fatal(err)
	}

	return func(c *gin.Context) {
		safeDB.mu.Lock()
		defer safeDB.mu.Unlock()

		// Bind JSON payload to UserSignup struct
		var user UserSignup
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// TODO: Hash the password

		// Insert the user into the database
		if _, err = stmt.Exec(user.Username, user.Password, user.Name, user.Animal); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"result": "success"})
	}
}

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
