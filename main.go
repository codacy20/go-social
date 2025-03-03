package main

import (
	"social/internal/infrastructure/http"
	"social/internal/service/comment"
	"social/routes"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Set up logger
	logger := log.New()
	logger.SetLevel(log.InfoLevel)

	// Create a new Gin router
	router := gin.Default()

	// Set up dependency injection for comments
	commentRepo := http.NewCommentRepository(logger)
	commentService := comment.NewCommentService(commentRepo, logger)
	commentHandler := http.NewCommentHandler(commentService, logger)

	logger.Info("Starting Social API server on port 8080")

	// Register routes
	routes.SetupReadyRoute(router)
	routes.SetupPostsRoute(router)
	commentHandler.RegisterRoutes(router)

	// Run the server
	if err := router.Run(":8080"); err != nil {
		logger.Fatalf("Failed to run server: %v", err)
	}
}
