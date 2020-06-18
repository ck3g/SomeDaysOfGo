package app

import "github.com/gin-gonic/gin"

var router = gin.Default()

// StartApp starts the main application
func StartApp() {
	mapUrls()

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
