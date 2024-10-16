package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/dtos/tandch"
	"server/interface/Service"
	"server/utils"
	"strconv"
)

type TAndChController struct {
	tAndChService Service.TAndChServiceInterface
}

func NewTAndChController(tAndChService Service.TAndChServiceInterface) *TAndChController {
	return &TAndChController{tAndChService: tAndChService}
}

func (e *TAndChController) CreateTAndCh(c *gin.Context) {
	var tAndChCreateDto tandch.CreateTandChDto
	if err := c.ShouldBind(&tAndChCreateDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	if err := tAndChCreateDto.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	tAndCh, err := e.tAndChService.CreateTAndCh(&tAndChCreateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Create tAndCh failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Create tAndCh successfully",
		Data:    tAndCh,
		Error:   "",
	})
}

func (e *TAndChController) DeleteTAndCh(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("tAndChId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid TAndCh Id",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	if err := e.tAndChService.DeleteTAndCh(id); err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Delete tAndCh failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Delete tAndCh successfully",
		Data:    nil,
		Error:   "",
	})
	return
}
func (e *TAndChController) UpdateTAndCh(c *gin.Context) {
	var tAndChUpdateDto tandch.UpdateTandChDto
	id, err := strconv.Atoi(c.Param("tAndChId"))

	if err := c.ShouldBind(&tAndChUpdateDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	tAndCh, err := e.tAndChService.UpdateTAndCh(id, tAndChUpdateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Update tAndCh failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Update tAndCh successfully",
		Data:    tAndCh,
		Error:   "",
	})
}

func (r *TAndChController) GetAllTAndCh(c *gin.Context) {
	tAndCh, err := r.tAndChService.GetAllTAndChs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Get tAndCh failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Get tAndCh successfully",
		Data:    tAndCh,
		Error:   "",
	})
	return
}
