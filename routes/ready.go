package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupReadyRoute(router *gin.Engine) {
	router.GET("/ready", readyHandler)
}

func readyHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Ready",
	})
}
