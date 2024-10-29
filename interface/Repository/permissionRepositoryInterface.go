package Repository

import (
	"server/dtos/permission"
	"server/models"
)

type PermissionRepositoryInterface interface {
	CreatePermission(createPermissionDto *permission.CreatePermissionDto) (*models.Permission, error)
	UpdatePermission(permissionId uint, dto permission.UpdatePermissionDto) (*models.Permission, error)
	DeletePermission(permissionId uint) error
	GetAllPermissions() ([]*models.Permission, error)
	GetPermissionById(permissionId uint) (*models.Permission, error)
}
