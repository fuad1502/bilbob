package middlewares

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var webappUrl = os.Getenv("PROTOCOL") + os.Getenv("HOSTNAME") + ":" + os.Getenv("WEBAPP_PORT")
var landingPageUrl = os.Getenv("PROTOCOL") + os.Getenv("HOSTNAME") + ":" + os.Getenv("LANDING_PAGE_PORT")

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		for _, e := range c.Errors {
			log.Println(e)
		}

		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get origins
		origins := c.Request.Header["Origin"]
		// Only set Access-Control-Allow-Origin to origin if origin either comes from landing-page webserver or webapp webserver
		var allowedOrigin string
		if len(origins) == 0 {
			allowedOrigin = ""
		} else {
			allowedOrigin = origins[0]
			log.Printf("Request Origin: %v\n", allowedOrigin)
			if allowedOrigin != webappUrl && allowedOrigin != landingPageUrl {
				allowedOrigin = ""
			}
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
