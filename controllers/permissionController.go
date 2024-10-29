package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/dtos/permission"
	"server/interface/Service"
	"server/utils"
	"strconv"
)

type PermissionController struct {
	permissionService Service.PermissionServiceInterface
}

func NewPermissionController(permissionService Service.PermissionServiceInterface) *PermissionController {
	return &PermissionController{
		permissionService: permissionService,
	}
}

func (r *PermissionController) CreatePermission(c *gin.Context) {
	var permissionCreateDto permission.CreatePermissionDto
	if err := c.ShouldBind(&permissionCreateDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	if err := permissionCreateDto.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	permission, err := r.permissionService.CreatePermission(&permissionCreateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Create permission failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, &utils.Response{
		Status:  http.StatusCreated,
		Message: "Create permission successfully",
		Data:    permission,
		Error:   "",
	})
}

func (r *PermissionController) DeletePermission(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("permissionId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	if err := r.permissionService.DeletePermission(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Delete permission failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Delete permission successfully",
		Data:    nil,
		Error:   "",
	})
	return
}

func (r *PermissionController) UpdatePermission(c *gin.Context) {
	var permissionUpdateDto permission.UpdatePermissionDto
	permissionId, err := strconv.Atoi(c.Param("permissionId"))

	if err := c.ShouldBind(&permissionUpdateDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	permission, err := r.permissionService.UpdatePermission(uint(permissionId), permissionUpdateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Update permission failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Update permission successfully",
		Data:    permission,
		Error:   "",
	})
}

func (r *PermissionController) GetAllPermission(c *gin.Context) {
	permissions, err := r.permissionService.GetAllPermissions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Get all permissions failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Get all permissions successfully",
		Data:    permissions,
		Error:   "",
	})
	return
}

func (r *PermissionController) GetPermissionById(c *gin.Context) {
	permissionId, err := strconv.ParseInt(c.Param("permissionId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	permission, err := r.permissionService.GetPermissionById(uint(permissionId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Permission get error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Permission get successfully",
		Data:    permission,
		Error:   "",
	})

}
