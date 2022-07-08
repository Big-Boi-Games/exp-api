package main

import (
	Db "exp-api/database"
	Mid "exp-api/middleware"
	Routes "exp-api/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/status", func(c *gin.Context) {
		fmt.Print("tester")
		c.JSON(200, gin.H{
			"Status": "All is well.",
		})
	})

	user := router.Group("/user")
	user.Use(Mid.AuthHandler())
	{
		Routes.Users(user)
	}

	Db.InitDatabase()

	router.Run()
}
