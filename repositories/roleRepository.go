package repositories

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"server/dtos/role"
	"server/interface/Repository"
	"server/models"
	"time"
)

type RoleRepository struct {
	DB *gorm.DB
}

func (r *RoleRepository) GetRoleById(roleId uint) (*models.Role, error) {
	var role models.Role
	if err := r.DB.Table("role").Where("id = ?", roleId).Preload("Permissions").First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepository) CreateRole(dto *role.CreateRoleDto) (*models.Role, error) {
	var permissions []models.Permission
	if len(dto.Permissions) > 0 {
		if err := r.DB.Where("id IN ?", dto.Permissions).Find(&permissions).Error; err != nil {
			log.Printf("Error finding permissions: %v", err)
			return nil, err
		}
	}

	newRole := &models.Role{
		RoleName:    dto.RoleName,
		Permissions: permissions,
	}

	if err := r.DB.Create(newRole).Error; err != nil {
		log.Printf("Error creating role: %v", err)
		return nil, err
	}

	return newRole, nil
}

func (r *RoleRepository) UpdateRole(roleID uint, dto role.UpdateRoleDto) (*models.Role, error) {
	var existingRole models.Role
	if err := r.DB.Preload("Permissions").First(&existingRole, roleID).Error; err != nil {
		log.Printf("Role not found: %v", err)
		return nil, err
	}

	var newPermissions []models.Permission
	if len(dto.Permissions) > 0 {
		if err := r.DB.Where("id IN ?", dto.Permissions).Find(&newPermissions).Error; err != nil {
			log.Printf("Error finding permissions: %v", err)
			return nil, err
		}
	}

	existingRole.RoleName = dto.RoleName

	if err := r.DB.Model(&existingRole).Association("Permissions").Replace(newPermissions); err != nil {
		log.Printf("Error updating role permissions: %v", err)
		return nil, err
	}

	if err := r.DB.Save(&existingRole).Error; err != nil {
		log.Printf("Error updating role: %v", err)
		return nil, err
	}

	return &existingRole, nil
}

func (r *RoleRepository) DeleteRole(roleId uint) error {
	result := r.DB.Table("role").Where("id = ?", roleId).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("role not found")
	}
	return nil
}

func (r *RoleRepository) GetAllRoles() ([]*models.Role, error) {
	var roles []*models.Role
	if err := r.DB.Preload("Permissions").Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func NewRoleRepository(db *gorm.DB) Repository.RoleRepositoryInterface {
	return &RoleRepository{
		DB: db,
	}
}
