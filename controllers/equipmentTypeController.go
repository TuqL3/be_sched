package controllers

import (
	"net/http"
	"server/dtos/equipmentType"
	"server/interface/Service"
	"server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EquipmentTypeController struct {
	EquipmentTypeService Service.EquipmentTypeServiceInterface
}

func NewEquipmentTypeController(EquipmentTypeService Service.EquipmentTypeServiceInterface) *EquipmentTypeController {
	return &EquipmentTypeController{
		EquipmentTypeService: EquipmentTypeService,
	}
}

func (e *EquipmentTypeController) CreateEquipmentType(c *gin.Context) {
	var EquipmentTypeCreateDto equipmentType.CreateEquipmentTypeDto
	if err := c.ShouldBind(&EquipmentTypeCreateDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	if err := EquipmentTypeCreateDto.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	EquipmentType, err := e.EquipmentTypeService.CreateEquipmentType(&EquipmentTypeCreateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Create EquipmentType failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Create EquipmentType successfully",
		Data:    EquipmentType,
		Error:   "",
	})
}

func (e *EquipmentTypeController) DeleteEquipmentType(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("equipmenttypeId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid EquipmentType Id",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	if err := e.EquipmentTypeService.DeleteEquipmentType(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Delete EquipmentType failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Delete EquipmentType successfully",
		Data:    nil,
		Error:   "",
	})
	return
}
func (e *EquipmentTypeController) UpdateEquipmentType(c *gin.Context) {
	var EquipmentTypeUpdateDto equipmentType.UpdateEquipmentTypeDto
	id, err := strconv.Atoi(c.Param("equipmenttypeId"))

	if err := c.ShouldBind(&EquipmentTypeUpdateDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	EquipmentType, err := e.EquipmentTypeService.UpdateEquipmentType(uint(id), EquipmentTypeUpdateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Update EquipmentType failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Update EquipmentType successfully",
		Data:    EquipmentType,
		Error:   "",
	})
}

func (r *EquipmentTypeController) GetAllEquipmentType(c *gin.Context) {
	EquipmentType, err := r.EquipmentTypeService.GetAllEquipmentTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Get EquipmentType failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Get EquipmentType successfully",
		Data:    EquipmentType,
		Error:   "",
	})
	return
}

func (r *EquipmentTypeController) GetEquipmentTypeById(c *gin.Context) {
	EquipmentTypeId, err := strconv.ParseInt(c.Param("equipmenttypeId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	EquipmentType, err := r.EquipmentTypeService.GetEquipmentTypeById(uint(EquipmentTypeId))
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
		Data:    EquipmentType,
		Error:   "",
	})

}
