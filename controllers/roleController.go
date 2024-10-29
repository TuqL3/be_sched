package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/dtos/role"
	"server/interface/Service"
	"server/utils"
	"strconv"
)

type RoleController struct {
	roleService Service.RoleServiceInterface
}

func NewRoleController(roleService Service.RoleServiceInterface) *RoleController {
	return &RoleController{
		roleService: roleService,
	}
}

func (r *RoleController) CreateRole(c *gin.Context) {
	var roleCreateDto role.CreateRoleDto
	if err := c.ShouldBind(&roleCreateDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	if err := roleCreateDto.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	role, err := r.roleService.CreateRole(&roleCreateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Create role failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, &utils.Response{
		Status:  http.StatusCreated,
		Message: "Create role successfully",
		Data:    role,
		Error:   "",
	})
}

func (r *RoleController) DeleteRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("roleId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	if err := r.roleService.DeleteRole(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Delete role failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Delete role successfully",
		Data:    nil,
		Error:   "",
	})
	return
}

func (r *RoleController) UpdateRole(c *gin.Context) {
	var roleUpdateDto role.UpdateRoleDto
	roleId, err := strconv.Atoi(c.Param("roleId"))

	if err := c.ShouldBind(&roleUpdateDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	role, err := r.roleService.UpdateRole(uint(roleId), roleUpdateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Update role failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Update role successfully",
		Data:    role,
		Error:   "",
	})
}

func (r *RoleController) GetAllRole(c *gin.Context) {
	roles, err := r.roleService.GetAllRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Get all roles failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Get all roles successfully",
		Data:    roles,
		Error:   "",
	})
	return
}

func (r *RoleController) GetRoleById(c *gin.Context) {
	roleId, err := strconv.ParseInt(c.Param("roleId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	role, err := r.roleService.GetRoleById(uint(roleId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Role get error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "Role get successfully",
		Data:    role,
		Error:   "",
	})

}
