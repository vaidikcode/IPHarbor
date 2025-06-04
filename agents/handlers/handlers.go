package handlers

import (
	"agents/models"
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Currently wokring for hugging face services
// Takes in model api_key(token in model) and provider
// The js backend service in running on local host. We forward to that.
// But originally the service endpoint model apikey provider should be on Story
func GenerateImage(c *gin.Context) {
	// new validation for the request body first
	payload := models.Generate{}
	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Originally we get endpoint from Story
	// But for now we are using a local service
	// For segreagating btw nebulus and huggingface he logic should be in services itself
	req, err := http.NewRequest("POST", "http://localhost:3000/huggingface/image", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	for key, values := range c.Request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"Failed to make request to service": err.Error()})
		return
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, gin.H{"Failed to read repsonse body": err.Error()})
		return
	}

	// Set the response headers
	for key, values := range resp.Header {
		for _, value := range values {
			c.Header(key, value)
		}
	}

	// Set the status code and return the response
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), respBody)
}

func GenerateText(c *gin.Context) {
	// new validation for the request body first
	payload := models.Generate{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Originally we get endpoint from Story
	// But for now we are using a local service
	req, err := http.NewRequest("POST", "http://localhost:3000/huggingface/text", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	for key, values := range c.Request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"Failed to make request to service": err.Error()})
		return
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, gin.H{"Failed to read repsonse body": err.Error()})
		return
	}

	// Set the response headers
	for key, values := range resp.Header {
		for _, value := range values {
			c.Header(key, value)
		}
	}

	// Set the status code and return the response
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), respBody)
}

func GenerateAudio(c *gin.Context) {
	// new validation for the request body first
	payload := models.Generate{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Originally we get endpoint from Story
	// But for now we are using a local service
	req, err := http.NewRequest("POST", "http://localhost:3000/huggingface/audio", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	for key, values := range c.Request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"Failed to make request to service": err.Error()})
		return
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, gin.H{"Failed to read repsonse body": err.Error()})
		return
	}

	// Set the response headers
	for key, values := range resp.Header {
		for _, value := range values {
			c.Header(key, value)
		}
	}

	// Set the status code and return the response
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), respBody)
}

func GenerateVideo(c *gin.Context) {
	// new validation for the request body first
	payload := models.Generate{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Originally we get endpoint from Story
	// But for now we are using a local service
	req, err := http.NewRequest("POST", "http://localhost:3000/huggingface/video", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	for key, values := range c.Request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"Failed to make request to service": err.Error()})
		return
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, gin.H{"Failed to read repsonse body": err.Error()})
		return
	}

	// Set the response headers
	for key, values := range resp.Header {
		for _, value := range values {
			c.Header(key, value)
		}
	}

	// Set the status code and return the response
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), respBody)
}
