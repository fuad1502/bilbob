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

type Following struct {
	Username string `json:"username" binding:"required"`
	State    string `json:"state" binding:"required"`
}

type Post struct {
	Username string    `json:"username" binding:"required"`
	PostText string    `json:"postText" binding:"required"`
	PostDate time.Time `json:"postDate"`
}

const session_timeout = 3600

func CreateLoginHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query("username")
		query := "SELECT password FROM Users WHERE username = $1"
		var saltAndHash string
		if err := safeDB.QueryRow(query, &saltAndHash, username); err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusOK, gin.H{"verified": false})
				return
			} else {
				c.Error(errors.New(err, c, "userLoginHandler"))
				return
			}
		}
		submittedPassword := c.Query("password")
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
}

func CreateGetUserHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId, err := c.Cookie("id")
		if err != nil || !sessions.IsLoggedIn(sessionId) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		requestedUser := c.Param("username")
		if user, err := safeDB.QueryGetUser(requestedUser); err == nil {
			c.JSON(http.StatusOK, user)
			return
		} else if err == dbs.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
			return
		} else {
			c.Error(errors.New(err, c, "userInfoHandler"))
			return
		}
	}
}

func CreateGetProfilePictureHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId, err := c.Cookie("id")
		if err != nil || !sessions.IsLoggedIn(sessionId) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		requestedUser := c.Param("username")
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

func CreatePutProfilePictureHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInAs, ok := getUsername(c)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		requestedUser := c.Param("username")
		if loggedInAs != requestedUser {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		file, err := c.FormFile("profile-picture")
		if err != nil {
			c.Error(errors.New(err, c, "CreatePutProfilePictureHandler"))
			return
		}
		path := "/assets/" + requestedUser + "/profile-picture/" + file.Filename
		if err := c.SaveUploadedFile(file, path); err != nil {
			c.Error(errors.New(err, c, "CreatePutProfilePictureHandler"))
			return
		}
		stmt := "UPDATE Users SET profile_pic_path = $1 WHERE username = $2"
		if err := safeDB.UpdateRow(stmt, path, requestedUser); err != nil {
			c.Error(errors.New(err, c, "CreatePutProfilePictureHandler"))
			return
		}
		c.JSON(http.StatusCreated, gin.H{"result": "success"})
	}
}

func CreateGetFollowingsHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId, err := c.Cookie("id")
		if err != nil || !sessions.IsLoggedIn(sessionId) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		requestedUser := c.Param("username")
		// TODO: handle only followers or self can see the followings of a user
		follows := c.Query("username")
		if follows != "" {
			query := `
			SELECT follows, state
			FROM Followings
			WHERE username = $1 AND follows = $2
			`
			var following Following
			if err := safeDB.QueryRow(query, &following, requestedUser, follows); err != nil {
				if err == sql.ErrNoRows {
					c.AbortWithStatus(http.StatusNotFound)
					return
				} else {
					c.Error(errors.New(err, c, "userFollowingsHandler"))
					return
				}
			}
			c.JSON(http.StatusOK, following)
		} else {
			query := `
			SELECT follows, state
			FROM Followings
			WHERE username = $1
			`
			followings := make([]Following, 0)
			if newFollowings, err := safeDB.Query(query, followings, requestedUser); err != nil {
				c.Error(errors.New(err, c, "userFollowingsHandler"))
				return
			} else {
				c.JSON(http.StatusOK, newFollowings)
				return
			}
		}
	}
}

func CreateGetFollowersHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId, err := c.Cookie("id")
		if err != nil || !sessions.IsLoggedIn(sessionId) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		requestedUser := c.Param("username")
		query := `
		SELECT username, state
		FROM Followings
		WHERE follows = $1
		`
		followers := make([]Following, 0)
		if newFollowers, err := safeDB.Query(query, followers, requestedUser); err != nil {
			c.Error(errors.New(err, c, "userFollowingsHandler"))
			return
		} else {
			c.JSON(http.StatusOK, newFollowers)
			return
		}
	}
}

func CreatePostFollowingHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInAs, ok := getUsername(c)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		username := c.Param("username")
		if username != loggedInAs {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var following Following
		if err := c.BindJSON(&following); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if following.State == "requested" {
			// TODO: For now, assume every profile is public, so every follow request is automatically accepted
			// Later, create profile privacy field in Users table and check that field first.
			following.State = "follows"
			stmt := `
			INSERT INTO Followings(username, follows, state) VALUES ($1, $2, $3)
			`
			row := struct {
				Username string
				Follows  string
				State    string
			}{Username: username, Follows: following.Username, State: following.State}
			if err := safeDB.InsertRow(stmt, &row); err != nil {
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

func CreateDeleteFollowingHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInAs, ok := getUsername(c)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		username := c.Param("username")
		if username != loggedInAs {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		followingUsername := c.Param("followingUsername")
		stmt := "DELETE FROM Followings WHERE username=$1 AND follows=$2"
		if err := safeDB.DeleteRow(stmt, username, followingUsername); err != nil {
			c.Error(errors.New(err, c, "CreateDeleteFollowingsHandler"))
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": "success"})
	}
}

func CreateGetUsersHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId, err := c.Cookie("id")
		if err != nil || !sessions.IsLoggedIn(sessionId) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		like := c.Query("like")
		re := regexp.MustCompile(`[^ \w_]`)
		like = re.ReplaceAllString(like, "")
		like = strings.ToLower(like)
		like = "%" + like + "%"
		query := `
		SELECT username, name, animal
		FROM Users
		WHERE LOWER(name) LIKE $1 OR LOWER(username) LIKE $1
		`
		users := make([]dbs.User, 0)
		if newUsers, err := safeDB.Query(query, users, like); err != nil {
			c.Error(errors.New(err, c, "CreateGetUsersHandler"))
			return
		} else {
			c.JSON(http.StatusOK, newUsers)
		}
	}
}

func CreatePostUserHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user UserSignup
		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if _, err := safeDB.QueryGetUser(user.Username); err == nil {
			c.Status(http.StatusOK)
			return
		} else if err != dbs.ErrNoRows {
			c.Error(errors.New(err, c, "createPostUserHandler"))
			return
		}
		saltAndHash, err := passwords.GenerateSaltAndHash(user.Password)
		if err != nil {
			c.Error(errors.New(err, c, "createPostUserHandler"))
			return
		}
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
		sessionId, err := c.Cookie("id")
		if err != nil || !sessions.IsLoggedIn(sessionId) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		requestedUser := c.Query("username")
		includeFollowings := c.Query("includeFollowings")
		var query string
		if includeFollowings == "true" {
			query = `
			SELECT T.username, T.post_text, T.post_date
			FROM (
			SELECT P.username, P.post_text, P.post_date
			FROM Posts AS P
			WHERE P.username = $1
			UNION
			SELECT P.username, P.post_text, P.post_date
			FROM Posts AS P JOIN Followings AS F
			ON P.username = F.follows
			WHERE F.username = $1) AS T
			ORDER BY T.post_date DESC
			`
		} else if includeFollowings == "false" {
			query = `
			SELECT username, post_text, post_date
			FROM Posts
			WHERE username = $1
			ORDER BY post_date DESC
			`
		} else {
			c.AbortWithStatus(http.StatusBadRequest)
		}
		posts := make([]Post, 0)
		if newPosts, err := safeDB.Query(query, posts, requestedUser); err != nil {
			c.Error(errors.New(err, c, "CreateGetPostsHandler"))
			return
		} else {
			c.JSON(http.StatusOK, newPosts.([]Post))
		}
	}
}

func CreatePostPostHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInAs, ok := getUsername(c)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var post Post
		if err := c.BindJSON(&post); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if loggedInAs != post.Username {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		post.PostDate = time.Now()
		stmt := `
		INSERT INTO Posts(username, post_text, post_date)
		VALUES ($1, $2, $3)
		`
		if err := safeDB.InsertRow(stmt, &post); err != nil {
			c.Error(errors.New(err, c, "CreatePostPostHandler"))
			return
		}
		c.JSON(http.StatusCreated, gin.H{"result": "success"})
	}
}

func CreateExistsHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query("username")
		if _, err := safeDB.QueryGetUser(username); err == nil {
			c.JSON(http.StatusOK, gin.H{"exists": true})
			return
		} else if err == dbs.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{"exists": false})
			return
		} else {
			c.Error(errors.New(err, c, "CreateExistsHandler"))
			return
		}
	}
}

func CreateAuthorizeHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, ok := getUsername(c)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.JSON(http.StatusOK, gin.H{"username": username})
	}
}

func CreateLogoutHandler(safeDB *dbs.SafeDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId, err := c.Cookie("id")
		if err != nil || !sessions.IsLoggedIn(sessionId) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		sessions.DeleteSession(sessionId)
		c.SetCookie("id", "", 0, "/", environ.Hostname, false, true)
		c.Status(http.StatusNoContent)
	}
}

func getUsername(c *gin.Context) (string, bool) {
	sessionid, err := c.Cookie("id")
	if err != nil {
		return "", false
	}
	return sessions.GetUsername(sessionid)
}
