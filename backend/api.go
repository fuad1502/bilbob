package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
		} else {
			log.Printf("Connection failed: %v\n", err)
		}
		time.Sleep(time.Second)
		log.Println("Retrying database connection...")
	}
	if !connected {
		log.Fatal(err)
	}
	log.Println("Connected to the database!")
	defer safeDB.Close()

	// Create a new router
	router := gin.Default()

	// Add CORS middleware
	router.Use(middlewares.CORSMiddleware())

	// Add error middleware
	router.Use(middlewares.ErrorMiddleware())

	// Add users route
	router.GET("/users/:username/:action", routes.CreateUserActionHandler(safeDB))
	router.GET("/users/:username", routes.CreateGetUserInfoHandler(safeDB))
	router.GET("/users/:username/profile-picture", routes.CreateGetUserPicHandler(safeDB))
	router.POST("/users/:username/profile-picture", routes.CreatePostUserPicHandler(safeDB))
	router.GET("/users", routes.CreateGetUsersHandler(safeDB))
	router.POST("/users", routes.CreatePostUserHandler(safeDB))

	// Add followings route
	router.GET("/followings/:username", routes.CreateGetFollowingsHandler(safeDB))
	router.POST("/followings/:username", routes.CreatePostFollowingsHandler(safeDB))
	router.DELETE("/followings/:username", routes.CreateDeleteFollowingsHandler(safeDB))

	// Add posts route
	router.GET("/posts", routes.CreateGetPostsHandler(safeDB))
	router.POST("/posts", routes.CreatePostPostHandler(safeDB))

	// Add authorization route
	router.GET("/authorize", routes.CreateAuthorizeHandler(safeDB))
	router.GET("/logout", routes.CreateLogoutHandler(safeDB))

	// Run the server
	log.Println("Web service running")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
