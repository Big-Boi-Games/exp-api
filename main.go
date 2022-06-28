package main

import (
	Mid "exp-api/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status": "All is well.",
		})
	})

	user := router.Group("/user")
	user.Use(Mid.AuthHandler())
	{
		user.GET("/status", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"Status": "Bingus.",
			})
		})
	}

	router.Run()
}
