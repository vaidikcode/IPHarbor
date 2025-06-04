package main

import (
	"github.com/gin-gonic/gin"
	"agents/handlers"
)

// THE GO BACKEND IS BASIALLY TO GET A REQUEST WITH PROMPT AND TOKEN AND THEN FURTHER SEND IT TO REQUIRED  AGENT SDK
func main() {
	r := gin.Default()

	// Routes for generating content

	// based on type of asset output

	// text to assets - they further divided to their required sdk's to interact with agents
	r.POST("/generateTestToImage", handlers.GenerateImage)
	r.POST("/generateTextToText", handlers.GenerateText)
	r.POST("/generateTextToAudio", handlers.GenerateAudio)
	r.POST("/generateTextToVideo", handlers.GenerateVideo)

	// image to assets
	

	r.Run(":8080")
}