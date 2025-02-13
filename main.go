package main

import (
	"social/routes"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Set the logging level (optional).
	log.SetLevel(log.InfoLevel)

	// Create a new Gin router.
	router := gin.Default()

	log.Info("Starting Social API server on port 8080")

	// Register routes.
	routes.SetupReadyRoute(router)
	routes.SetupPostsRoute(router)

	// Run the server.
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
