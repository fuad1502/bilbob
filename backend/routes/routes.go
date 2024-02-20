package routes

import (
	"database/sql"
	"encoding/hex"
	"log"
	"net/http"
	"time"

	"github.com/fuad1502/bilbob-backend/dbs"
	"github.com/fuad1502/bilbob-backend/errors"
	"github.com/fuad1502/bilbob-backend/passwords"
	"github.com/fuad1502/bilbob-backend/sessions"
	"github.com/gin-gonic/gin"
)

type UserSignup struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Animal   string `json:"animal" binding:"required"`
}

type UserInfo struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Animal   string `json:"animal"`
}

type Follows struct {
	Follows string `json:"follows"`
	State   string `json:"state"`
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
		sessionId := sessions.CreateSession(username)
		c.SetCookie("id", sessionId, 3600, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{"verified": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"verified": false})
	}
}

func userInfoHandler(safeDB *dbs.SafeDB, c *gin.Context, stmt *sql.Stmt) {
	// Check if logged in
	sessionId, err := c.Cookie("id")
	if err != nil || !sessions.IsLoggedIn(sessionId) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Extract queried username
	requestedUser := c.Param("username")

	// Query database
	row := stmt.QueryRow(requestedUser)
	var userInfo UserInfo
	if err := row.Scan(&userInfo.Username, &userInfo.Name, &userInfo.Animal); err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
			return
		} else {
			c.Error(errors.New(err, c, "userInfoHandler"))
			return
		}
	}

	// Return info
	c.JSON(http.StatusOK, userInfo)
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

	// Prepare SQL statement for getting user info
	stmt3, err := safeDB.DB.Prepare(`
		SELECT username, name, animal
		FROM Users
		WHERE username = $1
	`)
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
		} else if action == "info" {
			userInfoHandler(safeDB, c, stmt3)
		} else {
			c.AbortWithStatus(http.StatusBadRequest)
		}
	}
}

func CreateGetFollowsHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	// Prepare SQL statement to get the following state in relation to a certain user
	stmt, err := safeDB.DB.Prepare(`
		SELECT follows, state
		FROM Follows
		WHERE username = $1 AND follows = $2
	`)
	if err != nil {
		log.Fatal(err)
	}

	return func(c *gin.Context) {
		// Get the username
		username, ok := getUsername(c)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Check username filter, if exists
		requestedUser := c.Query("username")
		// TODO: handle the case with no filter
		if requestedUser != "" {
			if requestedUser == username {
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}

			// Query follows status
			var follows Follows
			row := stmt.QueryRow(username, requestedUser)
			if err := row.Scan(follows.Follows, follows.State); err != nil {
				if err == sql.ErrNoRows {
					c.AbortWithStatus(http.StatusNotFound)
					return
				} else {
					c.Error(errors.New(err, c, "userFollowsHandler"))
					return
				}
			}
			c.JSON(http.StatusOK, follows)
		} else {
			c.AbortWithStatus(http.StatusBadRequest)
			return
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
			c.AbortWithStatus(http.StatusBadRequest)
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

func CreateGetPostsHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	// Prepare SQL statement for getting all posts
	safeDB.Lock.Lock()
	defer safeDB.Lock.Unlock()
	stmt, err := safeDB.DB.Prepare(`
		SELECT P.username, P.post_text, P.post_date 
		FROM	(SELECT P.username, P.post_text, P.post_date
			FROM Posts P
			WHERE P.username = $1
			UNION
			SELECT P.username, P.post_text, P.post_date
			FROM Posts P, Follows F
			WHERE F.username = $1 AND P.username = F.follows) AS P
		ORDER BY P.post_date DESC
		`)
	if err != nil {
		log.Fatal(err)
	}

	return func(c *gin.Context) {
		// Get the username
		username, ok := getUsername(c)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Query the database for all posts
		safeDB.Lock.Lock()
		defer safeDB.Lock.Unlock()
		rows, err := stmt.Query(username)
		if err != nil {
			c.Error(errors.New(err, c, "CreateGetPostsHandler"))
			return
		}
		defer rows.Close()

		// Iterate through the rows and add them to the response
		posts := make([]gin.H, 0)
		for rows.Next() {
			var username, postText string
			var postDate time.Time
			if err := rows.Scan(&username, &postText, &postDate); err != nil {
				c.Error(errors.New(err, c, "CreateGetPostsHandler"))
				return
			}
			posts = append(posts, gin.H{"username": username, "postText": postText, "postDate": postDate})
		}

		c.JSON(http.StatusOK, posts)
	}
}

func CreatePostPostHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	// Prepare SQL statement for posting a post
	safeDB.Lock.Lock()
	defer safeDB.Lock.Unlock()
	stmt, err := safeDB.DB.Prepare(`
		INSERT INTO Posts(username, post_text, post_date)
		VALUES ($1, $2, $3)`)
	if err != nil {
		log.Fatal(err)
	}

	return func(c *gin.Context) {
		// Get the username
		username, ok := getUsername(c)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Get the post text from the JSON payload
		var payload struct {
			PostText string `json:"postText"`
		}
		c.BindJSON(&payload)

		// Post the post to the database
		safeDB.Lock.Lock()
		defer safeDB.Lock.Unlock()
		if _, err := stmt.Exec(username, payload.PostText, time.Now()); err != nil {
			c.Error(errors.New(err, c, "CreatePostPostHandler"))
			return
		}

		c.JSON(http.StatusCreated, gin.H{"result": "success"})
	}
}

func CreateAuthorizeHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the username
		username, ok := getUsername(c)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Return username
		c.JSON(http.StatusOK, gin.H{"username": username})
	}
}

func getUsername(c *gin.Context) (string, bool) {
	sessionid, err := c.Cookie("id")
	if err != nil {
		return "", false
	}
	return sessions.GetUsername(sessionid)
}
