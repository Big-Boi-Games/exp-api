package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func AuthHandler() gin.HandlerFunc {

	return func(c *gin.Context) {

		godotenv.Load(".env")
		tester := os.Getenv("AUTH_HEADER")
		token := c.Request.Header.Get("Authorisation")

		if len(token) == 0 || token != tester {
			c.Abort()
			c.JSON(430, gin.H{"Error": "You are not authorised."})
		}

	}
}
