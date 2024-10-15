package Service

import (
	"server/dtos/user"
	"server/models"
)

type UserServiceInterface interface {
	Register(userRegisterDto *user.UserRegister) (*models.User, error)
	FindUserByUsername(username string) (*models.User, error)
	UpdateUser(userId int, dto user.UserUpdateDto) (*models.User, error)
	DeleteUser(userId int) error
	GetAllUsers(fullName string) ([]*models.User, error)
}
