package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/dtos/report"
	"server/interface/Service"
	"server/utils"
	"strconv"
)

type ReportController struct {
	reportService Service.ReportServiceInterface
}

func NewReportController(reportService Service.ReportServiceInterface) *ReportController {
	return &ReportController{reportService: reportService}
}

func (r *ReportController) GetCountReportOfRoom(c *gin.Context) {
	counts, err := r.reportService.GetQuantityReportOfRoom()
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
			Error:   "Error",
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Data:    counts,
		Error:   "",
		Message: "Get successfully",
	})
}

func (r *ReportController) CreateReport(c *gin.Context) {
	var reportCreateDto report.CreateReportDto

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, &utils.Response{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
			Data:    nil,
			Error:   "User not found in context",
		})
		return
	}

	claimsMap, ok := user.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Server Error",
			Data:    nil,
			Error:   "Unable to parse user claims",
		})
		return
	}

	claims := utils.Claims{}

	if idFloat, ok := claimsMap["id"].(float64); ok {
		claims.ID = fmt.Sprintf("%.0f", idFloat)
	} else if idStr, ok := claimsMap["id"].(string); ok {
		claims.ID = idStr
	}

	if roleStr, ok := claimsMap["role"].(string); ok {
		claims.Role = roleStr
	}

	if claims.Role == "giang_vien" {
		reportCreateDto.Status = "pending"
	}
	reportCreateDto.Status = "in_progress"

	if err := c.ShouldBind(&reportCreateDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "invalid request body",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	if err := reportCreateDto.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "invalid request body",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	report, err := r.reportService.CreateReport(&reportCreateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    report,
		Error:   "",
	})

}

func (r *ReportController) DeleteReport(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("reportId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "invalid request body",
			Data:    nil,
		})
		return
	}
	if err := r.reportService.DeleteReport(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    nil,
		Error:   "",
	})
	return
}
func (r *ReportController) UpdateReport(c *gin.Context) {
	var reportUpdateDto report.UpdateReportDto
	reportId, err := strconv.Atoi(c.Param("reportId"))

	if err := c.ShouldBind(&reportUpdateDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "invalid request body",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	report, err := r.reportService.UpdateReport(uint(reportId), reportUpdateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    report,
		Error:   "",
	})

}

func (r *ReportController) GetAllReport(c *gin.Context) {
	report, err := r.reportService.GetAllReports()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    report,
		Error:   "",
	})
	return
}

func (r *ReportController) GetReportById(c *gin.Context) {
	reportId, err := strconv.ParseInt(c.Param("reportId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	report, err := r.reportService.GetReportById(uint(reportId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Report get error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Report get successfully",
		Data:    report,
		Error:   "",
	})

}
