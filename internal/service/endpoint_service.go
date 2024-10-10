package service

import (
	"github.com/Prajna1999/dataspace-be/internal/models"
	"github.com/Prajna1999/dataspace-be/internal/repository"
)

type EndpointService struct {
	repo *repository.EndpointRepository
}

func NewEndpointService(endpointRepo *repository.EndpointRepository) *EndpointService {
	return &EndpointService{repo: endpointRepo}
}

func (s *EndpointService) CreateEndpoint(endpoint *models.EndPoint) error {
	return s.repo.Create(endpoint)
}

func (s *EndpointService) GetEndpointByID(id uint) (*models.EndPoint, error) {
	return s.repo.GetEndpointByID(id)
}
