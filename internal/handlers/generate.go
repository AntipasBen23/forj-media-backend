package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"forj-media-demo-backend/internal/models"
	"forj-media-demo-backend/internal/openai"
)

func GenerateContent(c *gin.Context) {
	var req models.GenerateRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}

	if req.RawInput == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "rawInput is required"})
		return
	}

	rawOutput, err := openai.Generate(req.RawInput, req.Product, req.Audience, req.Tone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Parse the JSON the model returned
	var response models.GenerateResponse
	if err := json.Unmarshal([]byte(rawOutput), &response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Model returned invalid JSON"})
		return
	}

	c.JSON(http.StatusOK, response)
}
