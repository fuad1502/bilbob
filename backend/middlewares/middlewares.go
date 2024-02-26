package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/fuad1502/bilbob/backend/environ"
	"github.com/gin-gonic/gin"
)

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
			allowedOrigin = strings.SplitN(origins[0], ":", 2)[0]
			if allowedOrigin != environ.WebappUrl && allowedOrigin != environ.LandingPageUrl {
				allowedOrigin = ""
			}
			allowedOrigin = origins[0]
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin",
			allowedOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials",
			"true")
		c.Writer.Header().Set("Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods",
			"POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
