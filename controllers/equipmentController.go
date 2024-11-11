package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/dtos/equipment"
	"server/interface/Service"
	"server/utils"
	"strconv"
)

type EquipmentController struct {
	equipmentService Service.EquipmentServiceInterface
}

func NewEquipmentController(equipmentService Service.EquipmentServiceInterface) *EquipmentController {
	return &EquipmentController{equipmentService: equipmentService}
}

func (e *EquipmentController) CreateEquipment(c *gin.Context) {
	var equipmentCreateDto equipment.CreateEquipmentDto
	if err := c.ShouldBind(&equipmentCreateDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	if err := equipmentCreateDto.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	equipment, err := e.equipmentService.CreateEquipment(&equipmentCreateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Create equipment failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Create equipment successfully",
		Data:    equipment,
		Error:   "",
	})
}

func (e *EquipmentController) GetQuantityByStatus(c *gin.Context) {
	count, err := e.equipmentService.GetQuantityByStatus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Get equipment failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Get equipment successfully",
		Data:    count,
		Error:   "",
	})
}

func (e *EquipmentController) DeleteEquipment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("equipmentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid Equipment Id",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	if err := e.equipmentService.DeleteEquipment(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Delete equipment failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Delete equipment successfully",
		Data:    nil,
		Error:   "",
	})
	return
}
func (e *EquipmentController) UpdateEquipment(c *gin.Context) {
	var equipmentUpdateDto equipment.UpdateEquipmentDto
	id, err := strconv.Atoi(c.Param("equipmentId"))

	if err := c.ShouldBind(&equipmentUpdateDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	equipment, err := e.equipmentService.UpdateEquipment(uint(id), equipmentUpdateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Update equipment failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Update equipment successfully",
		Data:    equipment,
		Error:   "",
	})
}

func (r *EquipmentController) GetAllEquipment(c *gin.Context) {
	equipment, err := r.equipmentService.GetAllEquipments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Get equipment failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Get equipment successfully",
		Data:    equipment,
		Error:   "",
	})
	return
}

func (r *EquipmentController) GetEquipmentById(c *gin.Context) {
	equipmentId, err := strconv.ParseInt(c.Param("equipmentId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	equipment, err := r.equipmentService.GetEquipmentById(uint(equipmentId))
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
		Data:    equipment,
		Error:   "",
	})

}
