package repository

import (
	"github.com/Prajna1999/dataspace-be/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

// create a category
func (r *CategoryRepository) Create(category *models.Category) error {
	return r.DB.Create(category).Error
}

// get all categories
func (r *CategoryRepository) GetAllCategories() ([]*models.Category, error) {
	var categories []*models.Category
	err := r.DB.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}
