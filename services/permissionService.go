package services

import (
	"server/dtos/permission"
	"server/interface/Repository"
	"server/interface/Service"
	"server/models"
)

type PermissionService struct {
	permissionRepository Repository.PermissionRepositoryInterface
}

func (r *PermissionService) GetPermissionById(permissionId uint) (*models.Permission, error) {
	return r.permissionRepository.GetPermissionById(permissionId)
}

func (r *PermissionService) CreatePermission(createPermissionDto *permission.CreatePermissionDto) (*models.Permission, error) {
	return r.permissionRepository.CreatePermission(createPermissionDto)
}

func (r *PermissionService) UpdatePermission(permissionId uint, dto permission.UpdatePermissionDto) (*models.Permission, error) {
	return r.permissionRepository.UpdatePermission(permissionId, dto)
}

func (r *PermissionService) DeletePermission(permissionId uint) error {
	return r.permissionRepository.DeletePermission(permissionId)
}

func (r *PermissionService) GetAllPermissions() ([]*models.Permission, error) {
	return r.permissionRepository.GetAllPermissions()
}

func NewPermissionService(permissionRepo Repository.PermissionRepositoryInterface) Service.PermissionServiceInterface {
	return &PermissionService{
		permissionRepository: permissionRepo,
	}
}
