package routes

import (
	"github.com/gin-gonic/gin"
)

func Users(g *gin.RouterGroup) {
	g.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status": "Bingus.",
		})
	})
}
