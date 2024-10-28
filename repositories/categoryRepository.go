package repositories

import (
	"errors"
	"server/dtos/category"
	"server/interface/Repository"
	"server/models"
	"time"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func (c *CategoryRepository) CreateCategory(createCategoryDto *category.CreateCategoryDto) (*models.Category, error) {
	if err := c.DB.Table("categories").Create(createCategoryDto).Error; err != nil {
		return nil, err
	}
	m := &models.Category{
		Name: createCategoryDto.Name,
	}

	return m, nil
}

func (c *CategoryRepository) DeleteCategory(categoryId uint) error {
	result := c.DB.Table("categories").Where("id = ?", categoryId).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (c *CategoryRepository) GetAllCategorys() ([]*models.Category, error) {
	var categories []*models.Category
	if err := c.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *CategoryRepository) GetCategoryById(categoryId uint) (*models.Category, error) {
	var category models.Category
	if err := c.DB.Table("categories").Where("id = ?", categoryId).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c *CategoryRepository) UpdateCategory(categoryId uint, dto category.UpdateCategoryDto) (*models.Category, error) {
	var existingCategory models.Category
	if err := c.DB.Table("categories").Where("id = ?", categoryId).First(&existingCategory).Error; err != nil {
		return nil, err
	}
	updates := map[string]interface{}{
		"name": dto.Name,
	}
	if err := c.DB.Table("categories").Where("id = ?", categoryId).Updates(updates).Error; err != nil {
		return nil, err
	}
	if err := c.DB.First(&existingCategory, categoryId).Error; err != nil {
		return nil, err
	}
	return &existingCategory, nil
}

func NewCategoryRepository(db *gorm.DB) Repository.CategoryRepositoryInterface {
	return &CategoryRepository{
		DB: db,
	}
}
