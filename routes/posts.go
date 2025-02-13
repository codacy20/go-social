package routes

import (
	"encoding/json"
	"net/http"
	"social/providers"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// SetupPostsRoute registers the /posts endpoint.
func SetupPostsRoute(router *gin.Engine) {
	router.GET("/posts", postsHandler)
	log.Info("Registered /posts route")
}

func postsHandler(c *gin.Context) {
	log.Info("Handling /posts request")
	// Pass the request context to the provider.
	data, status, err := providers.GetPosts(c.Request.Context())
	if err != nil {
		log.Errorf("Error in provider GetPosts: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Marshal the typed response into JSON.
	posts, err := json.Marshal(data)
	if err != nil {
		log.Errorf("Error marshalling posts data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Infof("Successfully retrieved posts data with status %d", status)
	c.Data(status, "application/json", posts)
}
