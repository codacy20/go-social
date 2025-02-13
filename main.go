package main

import (
	"fmt"

	"social/routes" // import the routes package

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")
	router := gin.Default()

	routes.SetupReadyRoute(router)
	routes.SetupPostsRoute(router)

	router.Run(":8080")
}
