package Service

import (
	"server/dtos/user"
	"server/models"
)

type UserServiceInterface interface {
	Register(userRegisterDto *user.UserRegister) (*models.User, error)
	FindUserByUsername(username string) (*models.User, error)
	GetUserById(userId uint) (*models.User, error)
	UpdateUser(userId uint, dto user.UserUpdateDto) (*models.User, error)
	DeleteUser(userId uint) error
	GetAllUsers(fullName string) ([]*models.User, error)
}
