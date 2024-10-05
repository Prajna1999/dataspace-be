package routes

import (
	"net/http"

	"github.com/Prajna1999/dataspace-be/internal/service"
	"github.com/gin-gonic/gin"
)

type CategoryRoutes struct {
	categoryService *service.CategoryService
}

func NewCategoryRoutes(categoryService *service.CategoryService) *CategoryRoutes {
	return &CategoryRoutes{
		categoryService: categoryService,
	}
}

func (cat *CategoryRoutes) Setup(router *gin.RouterGroup) {
	cats := router.Group("/categories")
	{
		cats.POST("")
		cats.GET("")
	}
}

func (cat *CategoryRoutes) createCategory(c *gin.Context) {
	var input struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return

	}
	err := cat.categoryService.CreateCategory(input.Name, input.Description)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Catgory created successfully"})
}

func (cat *CategoryRoutes) getCategories(c *gin.Context) {
	cats, err := cat.categoryService.GetAllCategories()

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusFound, cats)
}
