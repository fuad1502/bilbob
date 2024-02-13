package routes

import (
	"database/sql"
	"encoding/hex"
	"log"
	"net/http"

	"github.com/fuad1502/bilbob-backend/dbs"
	"github.com/fuad1502/bilbob-backend/errors"
	"github.com/fuad1502/bilbob-backend/passwords"
	"github.com/gin-gonic/gin"
)

type UserSignup struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Animal   string `json:"animal" binding:"required"`
}

func userExistsHandler(safeDB *dbs.SafeDB, c *gin.Context, stmt *sql.Stmt) {
	// Get the username from the URL
	username := c.Param("username")

	// Query the database for the user
	safeDB.Lock.Lock()
	defer safeDB.Lock.Unlock()
	row := stmt.QueryRow(username)
	if err := row.Scan(&username); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{"exists": false})
			return
		} else {
			c.Error(errors.New(err, c, "userExistsHandler"))
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"exists": true})
}

func userLoginHandler(safeDB *dbs.SafeDB, c *gin.Context, stmt *sql.Stmt) {
	// Get the username from the URL
	username := c.Param("username")

	// Query the database for the user
	safeDB.Lock.Lock()
	defer safeDB.Lock.Unlock()
	var saltAndHash string
	row := stmt.QueryRow(username)
	if err := row.Scan(&saltAndHash); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{"verified": false})
			return
		} else {
			c.Error(errors.New(err, c, "userLoginHandler"))
			return
		}
	}

	// Get the submitted password from the URL
	submittedPassword := c.Query("password")

	// Hash password with salt and compare with stored hash
	// TODO: Encapsulate the following logic into a function
	salt := saltAndHash[:passwords.SaltSize*2]
	storedHash := saltAndHash[passwords.SaltSize*2:]
	saltBytes, err := hex.DecodeString(salt)
	if err != nil {
		c.Error(errors.New(err, c, "userLoginHandler"))
		return
	}
	computedHash := passwords.HashPassword(submittedPassword, saltBytes)
	if computedHash == storedHash {
		c.JSON(http.StatusOK, gin.H{"verified": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"verified": false})
	}
}

func CreateUserActionHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	// Prepare SQL statement for checking if a user exists
	safeDB.Lock.Lock()
	defer safeDB.Lock.Unlock()
	stmt1, err := safeDB.DB.Prepare("SELECT username FROM Users WHERE username = $1")
	if err != nil {
		log.Fatal(err)
	}

	// Prepare SQL statement for checking password hash
	stmt2, err := safeDB.DB.Prepare("SELECT password FROM Users WHERE username = $1")
	if err != nil {
		log.Fatal(err)
	}

	return func(c *gin.Context) {
		action := c.Param("action")
		// Call the appropriate handler based on the action
		if action == "exists" {
			userExistsHandler(safeDB, c, stmt1)
		} else if action == "login" {
			userLoginHandler(safeDB, c, stmt2)
		} else {
			c.AbortWithStatus(http.StatusBadRequest)
		}
	}
}

func CreatePostUserHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	// Prepare SQL statement for adding a user
	safeDB.Lock.Lock()
	defer safeDB.Lock.Unlock()
	stmt, err := safeDB.DB.Prepare("INSERT INTO Users (username, password, name, animal) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Fatal(err)
	}

	return func(c *gin.Context) {
		// Bind JSON payload to UserSignup struct
		var user UserSignup
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Hash the password
		salt, err := passwords.GenerateSalt()
		if err != nil {
			c.Error(errors.New(err, c, "createPostUserHandler"))
			return
		}
		hashedPassword := passwords.HashPassword(user.Password, salt)

		// Insert the user into the database
		safeDB.Lock.Lock()
		defer safeDB.Lock.Unlock()
		if _, err = stmt.Exec(user.Username, hex.EncodeToString(salt)+hashedPassword, user.Name, user.Animal); err != nil {
			c.Error(errors.New(err, c, "createPostUserHandler"))
			return
		}

		c.JSON(http.StatusCreated, gin.H{"result": "success"})
	}
}
