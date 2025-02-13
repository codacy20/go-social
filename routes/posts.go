package routes

import (
	"social/providers"

	"github.com/gin-gonic/gin"
)

// SetupPostsRoute registers the /posts endpoint.
func SetupPostsRoute(router *gin.Engine) {
	router.GET("/posts", postsHandler)
}

func postsHandler(c *gin.Context) {

	data, status, err := providers.GetPosts(c.Request.Context())
	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Data(status, "application/json", data)
}
