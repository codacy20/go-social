package routes

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// SetupReadyRoute registers the /ready endpoint.
func SetupReadyRoute(router *gin.Engine) {
	router.GET("/ready", readyHandler)
	log.Info("Registered /ready route")
}

func readyHandler(c *gin.Context) {
	log.Info("Handling /ready request")
	c.String(200, "Ready")
}
