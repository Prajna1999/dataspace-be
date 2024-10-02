package app

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Prajna1999/dataspace-be/internal/database"
	"github.com/Prajna1999/dataspace-be/internal/models"
	"github.com/Prajna1999/dataspace-be/internal/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	router  *gin.Engine
	db      *gorm.DB
	orgRepo *repository.OrganizationRepository
}

type HealthCheckResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// function that returns a pointet to the App
func NewApp() (*App, error) {
	db, err := database.InitDB()

	if err != nil {
		log.Fatalf("Failed to initialize the DB %v", err)
	}

	// Automigrate the org model
	err = db.AutoMigrate(&models.Organization{})
	if err != nil {
		log.Fatalf("Failed to migrate the Org model %v", err)
	}

	// declare and initialize tha app
	app := &App{
		router:  gin.Default(),
		db:      db,
		orgRepo: repository.NewOrganizationRepository(db),
	}
	app.setupRoutes()
	return app, nil
}

// setupRoutes function
func (a *App) setupRoutes() {
	v1 := a.router.Group("/api/v1")

	{
		v1.GET("/health-check", a.healthCheck)
		v1.POST("/org", a.createOrg)
		v1.GET("/org/:id", a.getOrg)
	}
}

func (a *App) Run() error {
	return a.router.Run(":8080")
}

func (a *App) healthCheck(c *gin.Context) {
	// check database connection
	var response HealthCheckResponse
	sqlDB, err := a.db.DB()
	if err != nil {
		response = HealthCheckResponse{
			Status:  "Error",
			Message: "Database Connection Error",
		}
		c.IndentedJSON(http.StatusInternalServerError, response)
		return
	}

	// ping the dabase
	err =
		sqlDB.Ping()
	if err != nil {
		response = HealthCheckResponse{
			Status:  "Error",
			Message: "Database Ping Failed",
		}
		c.IndentedJSON(http.StatusInternalServerError, response)
		return
	}

	response = HealthCheckResponse{
		Status:  "OK",
		Message: "All Sysetms Narmal 🚀",
	}

	c.IndentedJSON(http.StatusOK, response)
}

// route controllers

func (a *App) createOrg(c *gin.Context) {
	var org models.Organization

	if err := c.ShouldBindJSON(&org); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := a.orgRepo.Create(&org); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create organization"})
		return
	}

	c.IndentedJSON(http.StatusOK, org)
}
func (a *App) getOrg(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	org, err := a.orgRepo.GetByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, org)
}
