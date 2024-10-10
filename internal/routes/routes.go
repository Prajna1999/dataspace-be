package routes

import (
	"github.com/Prajna1999/dataspace-be/internal/service"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	orgRoutes      *OrganizationRoutes
	categoryRoutes *CategoryRoutes
	apiRoutes      *ApiRoutes
	//add other route groups here
}

func NewRoutes(
	orgService *service.OrganizationService,
	categoryService *service.CategoryService,
	apiService *service.ApiService,
) *Routes {
	return &Routes{
		orgRoutes:      NewOrganizationRoutes(orgService),
		categoryRoutes: NewCategoryRoutes(categoryService),
		apiRoutes:      NewApiRoutes(apiService),
		//add other routes here
	}
}

func (r *Routes) SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		r.orgRoutes.Setup(api)
		r.categoryRoutes.Setup(api)
		r.apiRoutes.Setup(api)
		//setup other routes
	}
}
