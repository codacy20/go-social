package routes

import (
	"encoding/json"
	"net/http"
	"social/providers"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// SetupCommentsRoute registers the /comments endpoints
func SetupCommentsRoute(router *gin.Engine) {
	router.GET("/comments", commentsHandler)
	router.GET("/posts/:postId/comments", postCommentsHandler)
	log.Info("Registered /comments routes")
}

func commentsHandler(c *gin.Context) {
	log.Info("Handling /comments request")
	data, status, err := providers.GetComments(c.Request.Context())
	if err != nil {
		log.Errorf("Error in provider GetComments: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	comments, err := json.Marshal(data)
	if err != nil {
		log.Errorf("Error marshalling comments data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Infof("Successfully retrieved comments data with status %d", status)
	c.Data(status, "application/json", comments)
}

func postCommentsHandler(c *gin.Context) {
	postID := c.Param("postId")
	log.Infof("Handling /posts/%s/comments request", postID)

	data, status, err := providers.GetCommentsByPostID(c.Request.Context(), postID)
	if err != nil {
		log.Errorf("Error in provider GetCommentsByPostID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	comments, err := json.Marshal(data)
	if err != nil {
		log.Errorf("Error marshalling comments data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Infof("Successfully retrieved comments for post %s with status %d", postID, status)
	c.Data(status, "application/json", comments)
}
