package Service

import (
	"server/dtos/user"
	"server/models"
)

type UserServiceInterface interface {
	Register(userRegisterDto *user.UserRegister) (*models.User, error)
	FindUserByUsername(username string) (*models.User, error)
	GetUserById(userId uint) (*models.User, error)
	UpdateUser(userId uint, dto user.UpdateUserDto, imageUrl string) (*models.User, error)
	DeleteUser(userId uint) error
	GetAllUsers(fullName, role string) ([]*models.User, error)
	GetCountUser() (int64, error)
}
