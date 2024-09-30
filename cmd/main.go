package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// define JSON object type
type HealthCheckResponse struct {
	Status string `json:"status"`
	// Message string `json:"message"`
}

func main() {
	router := gin.Default()

	// group the versioned API
	v1 := router.Group("/api/v1")

	{
		v1.GET("/health-check", healthCheck)
	}

	router.Run(":8080")
}

func healthCheck(c *gin.Context) {
	response := HealthCheckResponse{
		Status:  "OK",
		Message: "All Systems Narmal",
	}
	c.JSON(http.StatusOK, response)
}
