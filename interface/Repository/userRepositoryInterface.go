package Repository

import (
	"server/dtos/user"
	"server/models"
)

type UserRepositoryInterface interface {
	Register(userRegisterDto *user.UserRegister) (*models.User, error)
	FindUserByUsername(username string) (*models.User, error)
	GetUserById(userId uint) (*models.User, error)
	UpdateUser(userId uint, dto user.UpdateUserDto) (*models.User, error)
	DeleteUser(userId uint) error
	GetAllUsers(fullName string) ([]*models.User, error)
}
