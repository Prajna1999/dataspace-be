package routes

import (
	"github.com/Prajna1999/dataspace-be/internal/service"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	orgRoutes *OrganizationRoutes
	//add other route groups here
}

func NewRoutes(orgService *service.OrganizationService) *Routes {
	return &Routes{
		orgRoutes: NewOrganizationRoutes(orgService),
		//add other routes here
	}
}

func (r *Routes) SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		r.orgRoutes.Setup(api)
		//setup other routes
	}
}
