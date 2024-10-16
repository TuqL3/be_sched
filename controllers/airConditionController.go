package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/dtos/airCondition"
	"server/interface/Service"
	"server/utils"
	"strconv"
)

type AirConditionController struct {
	airConditionService Service.AirConditionServiceInterface
}

func NewAirConditionController(airConditionService Service.AirConditionServiceInterface) *AirConditionController {
	return &AirConditionController{airConditionService: airConditionService}
}

func (e *AirConditionController) CreateAirCondition(c *gin.Context) {
	var airConditionCreateDto aircondition.CreateAirConditionDto
	if err := c.ShouldBind(&airConditionCreateDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	if err := airConditionCreateDto.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	airCondition, err := e.airConditionService.CreateAirCondition(&airConditionCreateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Create airCondition failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Create airCondition successfully",
		Data:    airCondition,
		Error:   "",
	})
}

func (e *AirConditionController) DeleteAirCondition(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("airConditionId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid AirCondition Id",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	if err := e.airConditionService.DeleteAirCondition(id); err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Delete airCondition failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Delete airCondition successfully",
		Data:    nil,
		Error:   "",
	})
	return
}
func (e *AirConditionController) UpdateAirCondition(c *gin.Context) {
	var airConditionUpdateDto aircondition.UpdateAirConditionDto
	id, err := strconv.Atoi(c.Param("airConditionId"))

	if err := c.ShouldBind(&airConditionUpdateDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	airCondition, err := e.airConditionService.UpdateAirCondition(id, airConditionUpdateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Update airCondition failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Update airCondition successfully",
		Data:    airCondition,
		Error:   "",
	})
}

func (r *AirConditionController) GetAllAirCondition(c *gin.Context) {
	airCondition, err := r.airConditionService.GetAllAirConditions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Get airCondition failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Get airCondition successfully",
		Data:    airCondition,
		Error:   "",
	})
	return
}
