package routes

import (
	"net/http"
	"strconv"

	"github.com/Prajna1999/dataspace-be/internal/service"
	"github.com/gin-gonic/gin"
)

type EndPointRoutes struct {
	endpointService *service.EndpointService
}

func NewEndpointRoutes(endpointService *service.EndpointService) *EndPointRoutes {
	return &EndPointRoutes{endpointService: endpointService}
}

func (endpoint *EndPointRoutes) Setup(router *gin.RouterGroup) {
	endpoints := router.Group("/endpoints")
	{
		endpoints.GET("/:id", endpoint.getEndpointByID)
	}
}

func (h *EndPointRoutes) getEndpointByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	endpoint, err := h.endpointService.GetEndpointByID(uint(id))

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	if endpoint == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Endpoint does not exist for this id"})
	}
	c.IndentedJSON(http.StatusOK, endpoint)

}
