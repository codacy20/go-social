package http

import (
	"encoding/json"
	"net/http"
	"social/internal/domain/comment"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// CommentHandler handles HTTP requests for comments
type CommentHandler struct {
	service comment.Service
	logger  *log.Logger
}

// NewCommentHandler creates a new instance of CommentHandler
func NewCommentHandler(service comment.Service, logger *log.Logger) *CommentHandler {
	return &CommentHandler{
		service: service,
		logger:  logger,
	}
}

// RegisterRoutes registers the comment routes with the given router
func (h *CommentHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/comments", h.GetComments)
	router.GET("/posts/:postId/comments", h.GetCommentsByPostID)
	h.logger.Info("Registered /comments routes")
}

// GetComments handles the GET /comments endpoint
func (h *CommentHandler) GetComments(c *gin.Context) {
	h.logger.Info("Handler: Handling /comments request")
	data, status, err := h.service.GetComments(c.Request.Context())
	if err != nil {
		h.logger.Errorf("Handler: Error in service GetComments: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	comments, err := json.Marshal(data)
	if err != nil {
		h.logger.Errorf("Handler: Error marshalling comments data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.logger.Infof("Handler: Successfully retrieved comments data with status %d", status)
	c.Data(status, "application/json", comments)
}

// GetCommentsByPostID handles the GET /posts/:postId/comments endpoint
func (h *CommentHandler) GetCommentsByPostID(c *gin.Context) {
	postID := c.Param("postId")
	h.logger.Infof("Handler: Handling /posts/%s/comments request", postID)

	data, status, err := h.service.GetCommentsByPostID(c.Request.Context(), postID)
	if err != nil {
		h.logger.Errorf("Handler: Error in service GetCommentsByPostID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	comments, err := json.Marshal(data)
	if err != nil {
		h.logger.Errorf("Handler: Error marshalling comments data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.logger.Infof("Handler: Successfully retrieved comments for post %s with status %d", postID, status)
	c.Data(status, "application/json", comments)
}
