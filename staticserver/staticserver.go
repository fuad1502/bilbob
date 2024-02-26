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

	"github.com/fuad1502/bilbob/landing-page/handlers"
	ss "github.com/fuad1502/bilbob/staticserver/handlers"
)

func main() {
	log.SetPrefix("[Bilbob HTTP Server]: ")

	// Landing page handler
	landingPageHandler := handlers.LandingPageHandler
	// React page handler
	reactPageHandler := http.FileServer(http.Dir("/bin/build"))
	// Create main handler
	staticserverHandler := ss.New(landingPageHandler, reactPageHandler.ServeHTTP)
	// Run the server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: staticserverHandler,
	}
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()
	log.Println("Bilbob Web Server is running on port 8080! 🐱")

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
