package bloomapi

import "github.com/gin-gonic/gin"

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"ping": "pong",
	})
}
