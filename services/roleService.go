package services

import (
	"server/dtos/role"
	"server/interface/Repository"
	"server/interface/Service"
	"server/models"
)

type RoleService struct {
	roleRepository Repository.RoleRepositoryInterface
}

func (r *RoleService) GetRoleById(roleId uint) (*models.Role, error) {
	return r.roleRepository.GetRoleById(roleId)
}

func (r *RoleService) CreateRole(createRoleDto *role.CreateRoleDto) (*models.Role, error) {
	return r.roleRepository.CreateRole(createRoleDto)
}

func (r *RoleService) UpdateRole(roleId uint, dto role.UpdateRoleDto) (*models.Role, error) {
	return r.roleRepository.UpdateRole(roleId, dto)
}

func (r *RoleService) DeleteRole(roleId uint) error {
	return r.roleRepository.DeleteRole(roleId)
}

func (r *RoleService) GetAllRoles() ([]*models.Role, error) {
	return r.roleRepository.GetAllRoles()
}

func NewRoleService(roleRepo Repository.RoleRepositoryInterface) Service.RoleServiceInterface {
	return &RoleService{
		roleRepository: roleRepo,
	}
}
