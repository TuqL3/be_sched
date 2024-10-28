package controllers

import (
	"net/http"
	"server/dtos/category"
	"server/interface/Service"
	"server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService Service.CategoryServiceInterface
}

func NewCategoryController(categoryService Service.CategoryServiceInterface) *CategoryController {
	return &CategoryController{
		categoryService: categoryService,
	}
}

func (e *CategoryController) CreateCategory(c *gin.Context) {
	var categoryCreateDto category.CreateCategoryDto
	if err := c.ShouldBind(&categoryCreateDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	if err := categoryCreateDto.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	category, err := e.categoryService.CreateCategory(&categoryCreateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Create category failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Create category successfully",
		Data:    category,
		Error:   "",
	})
}

func (e *CategoryController) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid Category Id",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	if err := e.categoryService.DeleteCategory(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Delete category failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Delete category successfully",
		Data:    nil,
		Error:   "",
	})
	return
}
func (e *CategoryController) UpdateCategory(c *gin.Context) {
	var categoryUpdateDto category.UpdateCategoryDto
	id, err := strconv.Atoi(c.Param("categoryId"))

	if err := c.ShouldBind(&categoryUpdateDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	category, err := e.categoryService.UpdateCategory(uint(id), categoryUpdateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Update category failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Update category successfully",
		Data:    category,
		Error:   "",
	})
}

func (r *CategoryController) GetAllCategory(c *gin.Context) {
	category, err := r.categoryService.GetAllCategorys()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Get category failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Get category successfully",
		Data:    category,
		Error:   "",
	})
	return
}

func (r *CategoryController) GetCategoryById(c *gin.Context) {
	categoryId, err := strconv.ParseInt(c.Param("categoryId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	category, err := r.categoryService.GetCategoryById(uint(categoryId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Air condition get error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Air condition get successfully",
		Data:    category,
		Error:   "",
	})

}
