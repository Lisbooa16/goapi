package router

import "github.com/gin-gonic/gin"

func Initializer() {
	// Initializer router
	router := gin.Default()

	// Initialize routes
	InitializeRoutes(router)

	// run the server
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
