package routes

import (
	"database/sql"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/fuad1502/bilbob/backend/dbs"
	"github.com/fuad1502/bilbob/backend/environ"
	"github.com/fuad1502/bilbob/backend/errors"
	"github.com/fuad1502/bilbob/backend/passwords"
	"github.com/fuad1502/bilbob/backend/sessions"
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

type FollowingInfo struct {
	Username string `json:"username"`
	Follows  string `json:"follows"`
	State    string `json:"state"`
}

type Post struct {
	Username string    `json:"username"`
	PostText string    `json:"postText"`
	PostDate time.Time `json:"postDate"`
}

const session_timeout = 3600

func userExistsHandler(safeDB *dbs.SafeDB, c *gin.Context) {
	// Get the username from the URL
	username := c.Param("username")

	// Query the database for the user
	query := "SELECT username FROM Users WHERE username = $1"
	if err := safeDB.QueryRow(query, &username, username); err != nil {
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

func userLoginHandler(safeDB *dbs.SafeDB, c *gin.Context) {
	// Get the username from the URL
	username := c.Param("username")

	// Query the database for the user's salt & hash
	query := "SELECT password FROM Users WHERE username = $1"
	var saltAndHash string
	if err := safeDB.QueryRow(query, &saltAndHash, username); err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
			return
		} else {
			c.Error(errors.New(err, c, "userLoginHandler"))
			return
		}
	}

	// Get the submitted password from the URL
	submittedPassword := c.Query("password")

	// Verifiy the submitted password
	verified, err := passwords.VerifyPassword(submittedPassword, saltAndHash)
	if err != nil {
		c.Error(errors.New(err, c, "userLoginHandler"))
		return
	}
	if verified {
		sessionId := sessions.CreateSession(username)
		c.SetCookie("id", sessionId, session_timeout, "/", environ.Hostname, false, true)
		c.JSON(http.StatusOK, gin.H{"verified": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"verified": false})
	}
}

func CreateGetUserInfoHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if logged in
		sessionId, err := c.Cookie("id")
		if err != nil || !sessions.IsLoggedIn(sessionId) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Extract queried username
		requestedUser := c.Param("username")

		// Query database
		query := `
		SELECT username, name, animal
		FROM Users
		WHERE username = $1
		`
		var userInfo UserInfo
		if err := safeDB.QueryRow(query, &userInfo, requestedUser); err != nil {
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
}

func CreateGetUserPicHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if logged in
		sessionId, err := c.Cookie("id")
		if err != nil || !sessions.IsLoggedIn(sessionId) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Extract queried username
		requestedUser := c.Param("username")

		// Query database
		query := `
		SELECT profile_pic_path
		FROM Users
		WHERE username = $1
		`
		var path any
		if err := safeDB.QueryRow(query, &path, requestedUser); err != nil {
			if err == sql.ErrNoRows {
				c.AbortWithStatus(http.StatusNotFound)
				return
			} else {
				c.Error(errors.New(err, c, "CreateGetUserPicHandler"))
				return
			}
		}
		if path == nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.File(path.(string))
	}
}

func CreatePostUserPicHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the logged in user
		loggedInAs, ok := getUsername(c)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// Get the username
		username := c.Param("username")
		// Only the logged in user can update its profile picture
		if loggedInAs != username {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Get the picture
		file, err := c.FormFile("profile-picture")
		if err != nil {
			c.Error(errors.New(err, c, "CreatePostUserPicHandler"))
			return
		}

		// Save the picture
		path := "/assets/" + username + "/profile-picture/" + file.Filename
		if err := c.SaveUploadedFile(file, path); err != nil {
			c.Error(errors.New(err, c, "CreatePostUserPicHandler"))
			return
		}

		// Save the path to DB
		stmt := "UPDATE Users SET profile_pic_path = $1 WHERE username = $2"
		if err := safeDB.UpdateRow(stmt, path, username); err != nil {
			c.Error(errors.New(err, c, "CreatePostUserPicHandler"))
			return
		}
		c.JSON(http.StatusCreated, gin.H{"result": "success"})
	}
}

func CreateUserActionHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		action := c.Param("action")
		// Call the appropriate handler based on the action
		if action == "exists" {
			userExistsHandler(safeDB, c)
		} else if action == "login" {
			userLoginHandler(safeDB, c)
		} else {
			c.AbortWithStatus(http.StatusBadRequest)
		}
	}
}

func CreateGetFollowingsHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if logged in
		sessionId, err := c.Cookie("id")
		if err != nil || !sessions.IsLoggedIn(sessionId) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Get the username
		requestedUser := c.Param("username")
		// TODO: handle only followers or self can see the followings of a user

		// Check follows filter, if exists
		follows := c.Query("follows")
		// TODO: handle the case with no filter
		if follows != "" {
			// Query follows status
			query := `
			SELECT username, follows, state
			FROM Followings
			WHERE username = $1 AND follows = $2
			`
			var followingInfo FollowingInfo
			if err := safeDB.QueryRow(query, &followingInfo, requestedUser, follows); err != nil {
				if err == sql.ErrNoRows {
					c.AbortWithStatus(http.StatusNotFound)
					return
				} else {
					c.Error(errors.New(err, c, "userFollowingsHandler"))
					return
				}
			}
			c.JSON(http.StatusOK, followingInfo)
		} else {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}
}

func CreatePostFollowingsHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the logged in user
		loggedInAs, ok := getUsername(c)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Get username to POST to
		username := c.Param("username")
		if username != loggedInAs {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Get follows filter
		follows := c.Query("follows")
		if follows == "" {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// Get payload
		var payload FollowingInfo
		if err := c.BindJSON(&payload); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if follows != payload.Follows {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if payload.State == "requested" {
			// TODO: For now, assume every profile is public, so every follow request is automatically accepted
			// Later, create profile privacy field in Users table and check that field first.
			payload.State = "follows"
			stmt := `
			INSERT INTO Followings(username, follows, state) VALUES ($1, $2, $3)
			`
			if err := safeDB.InsertRow(stmt, &payload); err != nil {
				c.Error(errors.New(err, c, "CreatePostFollowingsHandler"))
				return
			}
		} else {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusCreated, gin.H{"result": "success"})
	}
}

func CreateDeleteFollowingsHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the logged in username
		loggedInAs, ok := getUsername(c)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Get username field selection
		username := c.Param("username")
		if username != loggedInAs {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Get follows field selection
		follows := c.Query("follows")
		if follows == "" {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		stmt := "DELETE FROM Followings WHERE username=$1 AND follows=$2"
		if err := safeDB.DeleteRow(stmt, username, follows); err != nil {
			c.Error(errors.New(err, c, "CreateDeleteFollowingsHandler"))
			return
		}

		c.JSON(http.StatusOK, gin.H{"result": "success"})
	}
}

func CreateGetUsersHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if logged in
		sessionId, err := c.Cookie("id")
		if err != nil || !sessions.IsLoggedIn(sessionId) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Check if has "like" filter
		like := c.Query("like")
		re := regexp.MustCompile(`[^ \w_]`)
		like = re.ReplaceAllString(like, "")
		if like == "" {
			c.JSON(http.StatusOK, gin.H{})
		}
		like = strings.ToLower(like)
		like = "%" + like + "%"

		// Find all users that has a name or username like "like"
		query := `
		SELECT username, name, animal
		FROM Users
		WHERE LOWER(name) LIKE $1 OR LOWER(username) LIKE $1
		`
		userInfos := make([]UserInfo, 0)
		if newUserInfos, err := safeDB.Query(query, userInfos, like); err != nil {
			c.Error(errors.New(err, c, "CreateGetUsersHandler"))
			return
		} else {
			c.JSON(http.StatusOK, newUserInfos)
		}
	}
}

func CreatePostUserHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bind JSON payload to UserSignup struct
		var user UserSignup
		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// Generate salt and hash from the password
		saltAndHash, err := passwords.GenerateSaltAndHash(user.Password)
		if err != nil {
			c.Error(errors.New(err, c, "createPostUserHandler"))
			return
		}

		// Insert the user into the database
		stmt := "INSERT INTO Users (username, password, name, animal) VALUES ($1, $2, $3, $4)"
		user.Password = saltAndHash
		if err = safeDB.InsertRow(stmt, &user); err != nil {
			c.Error(errors.New(err, c, "createPostUserHandler"))
			return
		}

		c.JSON(http.StatusCreated, gin.H{"result": "success"})
	}
}

func CreateGetPostsHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the username
		username, ok := getUsername(c)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Query the database for all posts
		query := `
		SELECT P.username, P.post_text, P.post_date 
		FROM	(SELECT P.username, P.post_text, P.post_date
			FROM Posts P
			WHERE P.username = $1
			UNION
			SELECT P.username, P.post_text, P.post_date
			FROM Posts P, Followings F
			WHERE F.username = $1 AND P.username = F.follows) AS P
		ORDER BY P.post_date DESC
		`
		posts := make([]Post, 0)
		if newPosts, err := safeDB.Query(query, posts, username); err != nil {
			c.Error(errors.New(err, c, "CreateGetPostsHandler"))
			return
		} else {
			c.JSON(http.StatusOK, newPosts.([]Post))
		}
	}
}

func CreatePostPostHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
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
		stmt := `
		INSERT INTO Posts(username, post_text, post_date)
		VALUES ($1, $2, $3)
		`
		post := Post{username, payload.PostText, time.Now()}
		if err := safeDB.InsertRow(stmt, &post); err != nil {
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

func CreateLogoutHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the username
		username, ok := getUsername(c)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Essentially delete the cookie
		c.SetCookie("id", "", 0, "/", environ.Hostname, false, true)

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
