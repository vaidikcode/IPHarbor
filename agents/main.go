package main

import (
	"github.com/gin-gonic/gin"
	"agents/handlers"
)

func main() {
	r := gin.Default()

	// Routes for generating content

	// based on type of asset output

	// text to assets
	r.POST("/generateTestToImage", handlers.GenerateImage)
	r.POST("/generateTextToText", handlers.GenerateText)
	r.POST("/generateTextToAudio", handlers.GenerateAudio)
	r.POST("/generateTextToVideo", handlers.GenerateVideo)

	// image to assets
	

	r.Run(":8080")
}