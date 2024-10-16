package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/dtos/computer"
	"server/interface/Service"
	"server/utils"
	"strconv"
)

type ComputerController struct {
	computerService Service.ComputerServiceInterface
}

func NewComputerController(computerService Service.ComputerServiceInterface) *ComputerController {
	return &ComputerController{computerService: computerService}
}

func (e *ComputerController) CreateComputer(c *gin.Context) {
	var computerCreateDto computer.CreateComputerDto
	if err := c.ShouldBind(&computerCreateDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	if err := computerCreateDto.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	computer, err := e.computerService.CreateCompute(&computerCreateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Create computer failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Create computer successfully",
		Data:    computer,
		Error:   "",
	})
}

func (e *ComputerController) DeleteComputer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("computerId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid Computer Id",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	if err := e.computerService.DeleteCompute(id); err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Delete computer failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Delete computer successfully",
		Data:    nil,
		Error:   "",
	})
	return
}
func (e *ComputerController) UpdateComputer(c *gin.Context) {
	var computerUpdateDto computer.UpdateComputerDto
	id, err := strconv.Atoi(c.Param("computerId"))

	if err := c.ShouldBind(&computerUpdateDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	computer, err := e.computerService.UpdateCompute(id, computerUpdateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Update computer failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Update computer successfully",
		Data:    computer,
		Error:   "",
	})
}

func (r *ComputerController) GetAllComputer(c *gin.Context) {
	computer, err := r.computerService.GetAllComputes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Get computer failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Get computer successfully",
		Data:    computer,
		Error:   "",
	})
	return
}
