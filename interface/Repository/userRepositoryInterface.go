package Repository

import (
	"server/dtos/user"
	"server/models"
)

type UserRepositoryInterface interface {
	Register(userRegisterDto *user.UserRegister) (*models.User, error)
	FindUserByUsername(username string) (*models.User, error)
	UpdateUser(userId int, dto user.UserUpdateDto) (*models.User, error)
	DeleteUser(userId int) error
	GetAllUsers() ([]*models.User, error)
}
