package services

import (
	"server/dtos/category"
	"server/interface/Repository"
	"server/interface/Service"
	"server/models"
)

type CategoryService struct {
	categoryRepository Repository.CategoryRepositoryInterface
}

func (c *CategoryService) CreateCategory(createCategoryDto *category.CreateCategoryDto) (*models.Category, error) {
	return c.categoryRepository.CreateCategory(createCategoryDto)
}

func (c *CategoryService) DeleteCategory(categoryId uint) error {
	return c.categoryRepository.DeleteCategory(categoryId)
}

func (c *CategoryService) GetAllCategorys() ([]*models.Category, error) {
	return c.categoryRepository.GetAllCategorys()
}

func (c *CategoryService) GetCategoryById(categoryId uint) (*models.Category, error) {
	return c.categoryRepository.GetCategoryById(categoryId)
}

func (c *CategoryService) UpdateCategory(categoryId uint, dto category.UpdateCategoryDto) (*models.Category, error) {
	return c.categoryRepository.UpdateCategory(categoryId, dto)
}

func NewCategoryService(categoryRepo Repository.CategoryRepositoryInterface) Service.CategoryServiceInterface {
	return &CategoryService{
		categoryRepository: categoryRepo,
	}
}
