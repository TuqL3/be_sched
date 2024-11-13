package repositories

import (
	"errors"
	"gorm.io/gorm"
	"server/dtos/permission"
	"server/interface/Repository"
	"server/models"
	"time"
)

type PermissionRepository struct {
	DB *gorm.DB
}

func (r *PermissionRepository) GetPermissionById(permissionId uint) (*models.Permission, error) {
	var permission models.Permission
	if err := r.DB.Table("permission").Where("id = ?", permissionId).First(&permission).Error; err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *PermissionRepository) CreatePermission(createPermissionDto *permission.CreatePermissionDto) (*models.Permission, error) {
	if err := r.DB.Table("permission").Create(createPermissionDto).Error; err != nil {
		return nil, err
	}

	m := &models.Permission{
		PermissionName: createPermissionDto.PermissionName,
	}
	return m, nil
}

func (r *PermissionRepository) UpdatePermission(permissionId uint, dto permission.UpdatePermissionDto) (*models.Permission, error) {
	var existingPermission models.Permission
	if err := r.DB.Table("permission").Where("id = ?", permissionId).First(&existingPermission).Error; err != nil {
		return nil, err
	}
	updates := map[string]interface{}{
		"permission_name": dto.PermissionName,
	}
	if err := r.DB.Table("permission").Where("id = ?", permissionId).Updates(updates).Error; err != nil {
		return nil, err
	}
	if err := r.DB.First(&existingPermission, permissionId).Error; err != nil {
		return nil, err
	}
	return &existingPermission, nil
}

func (r *PermissionRepository) DeletePermission(permissionId uint) error {
	result := r.DB.Table("permission").Where("id = ?", permissionId).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("permission not found")
	}
	return nil
}

func (r *PermissionRepository) GetAllPermissions() ([]*models.Permission, error) {
	var permission []*models.Permission
	if err := r.DB.Find(&permission).Error; err != nil {
		return nil, err
	}
	return permission, nil
}

func NewPermissionRepository(db *gorm.DB) Repository.PermissionRepositoryInterface {
	return &PermissionRepository{
		DB: db,
	}
}
