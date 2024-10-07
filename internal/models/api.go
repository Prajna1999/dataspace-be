package models

import "gorm.io/gorm"

type Api struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryID  uint
}
