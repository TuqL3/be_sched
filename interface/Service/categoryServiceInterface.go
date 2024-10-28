package Service

import (
	"server/dtos/category"
	"server/models"
)

type CategoryServiceInterface interface {
	CreateCategory(createCategoryDto *category.CreateCategoryDto) (*models.Category, error)
	UpdateCategory(categoryId uint, dto category.UpdateCategoryDto) (*models.Category, error)
	DeleteCategory(categoryId uint) error
	GetAllCategorys() ([]*models.Category, error)
	GetCategoryById(categoryId uint) (*models.Category, error)
}
