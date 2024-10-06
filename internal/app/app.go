package app

import (
	"log"
	"net/http"

	"github.com/Prajna1999/dataspace-be/internal/database"
	"github.com/Prajna1999/dataspace-be/internal/models"
	"github.com/Prajna1999/dataspace-be/internal/repository"
	"github.com/Prajna1999/dataspace-be/internal/routes"
	"github.com/Prajna1999/dataspace-be/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	router          *gin.Engine
	db              *gorm.DB
	orgService      *service.OrganizationService
	categoryService *service.CategoryService
	routes          *routes.Routes
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
	err = db.AutoMigrate(&models.Organization{}, &models.Category{})
	if err != nil {
		log.Fatalf("Failed to migrate the Org model %v", err)
	}

	orgRepo := repository.NewOrganizationRepository(db)
	orgService := service.NewOragnizationService(orgRepo)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	// declare and initialize tha app
	app := &App{
		router:          gin.Default(),
		db:              db,
		orgService:      orgService,
		categoryService: categoryService,
		routes:          routes.NewRoutes(orgService, categoryService),
	}
	app.setupRoutes()
	return app, nil
}

// setupRoutes function
func (a *App) setupRoutes() {
	a.router.GET("/api/v1/health-check", a.healthCheck)
	a.routes.SetupRoutes(a.router)
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
