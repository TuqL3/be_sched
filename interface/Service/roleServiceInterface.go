package Service

import (
	"server/dtos/role"
	"server/models"
)

type RoleServiceInterface interface {
	CreateRole(createRoleDto *role.CreateRoleDto) (*models.Role, error)
	UpdateRole(roleId uint, dto role.UpdateRoleDto) (*models.Role, error)
	DeleteRole(roleId uint) error
	GetAllRoles() ([]*models.Role, error)
	GetRoleById(roleId uint) (*models.Role, error)
}
