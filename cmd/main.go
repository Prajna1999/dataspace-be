package main

import (
	"log"
	"net/http"

	"github.com/Prajna1999/dataspace-be/internal/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// define JSON object type
type HealthCheckResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var db *gorm.DB

func main() {
	var err error

	db, err = database.InitDB()

	if err != nil {
		log.Fatalf("Failed to initialize the DB %v", err)
	}
	router := gin.Default()

	// group the versioned API
	v1 := router.Group("/api/v1")
	{
		v1.GET("/health-check", healthCheck)
	}
	router.Run(":8080")

}

func healthCheck(c *gin.Context) {
	// check the database connection
	sqlDB, err := db.DB()
	var response HealthCheckResponse
	if err != nil {
		response = HealthCheckResponse{
			Status:  "Error",
			Message: "Database Connection error",
		}

		c.IndentedJSON(http.StatusInternalServerError, response)
		return
	}

	// ping the database
	err = sqlDB.Ping()
	if err != nil {
		response = HealthCheckResponse{
			Status:  "Error",
			Message: "Database ping failed",
		}
		c.IndentedJSON(http.StatusInternalServerError, response)
		return
	}

	response = HealthCheckResponse{
		Status:  "Ok",
		Message: "All Systems Narmal 🚀",
	}
	log.Printf("Connected to neon db")
	c.IndentedJSON(http.StatusOK, response)
}
