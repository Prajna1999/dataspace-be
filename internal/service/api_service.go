package service

import (
	"github.com/Prajna1999/dataspace-be/internal/models"
	"github.com/Prajna1999/dataspace-be/internal/repository"
)

type ApiService struct {
	repo *repository.ApiRepository
}

func NewApiService(repo *repository.ApiRepository) *ApiService {
	return &ApiService{repo: repo}
}

// create a new api
func (s *ApiService) CreateApi(api *models.Api) error {
	return s.repo.Create(api)
}

// get all apis or filtered
func (s *ApiService) GetAllApis(filters map[string]interface{}) ([]*models.Api, error) {
	return s.repo.GetAll(filters)
}
