package services

import (
	"server/dtos/user"
	"server/interface/Repository"
	"server/interface/Service"
	"server/models"
)

type UserService struct {
	userRepository Repository.UserRepositoryInterface
}

func (u *UserService) GetAllUsers() ([]*models.User, error) {
	return u.userRepository.GetAllUsers()
}

func (u *UserService) DeleteUser(userId int) error {
	return u.userRepository.DeleteUser(userId)
}

func (u *UserService) UpdateUser(userId int, dto user.UserUpdateDto) (*models.User, error) {
	return u.userRepository.UpdateUser(userId, dto)
}

func (u *UserService) FindUserByUsername(username string) (*models.User, error) {
	return u.userRepository.FindUserByUsername(username)
}

func (u *UserService) Register(userRegisterDto *user.UserRegister) (*models.User, error) {
	return u.userRepository.Register(userRegisterDto)
}

func NewUserService(userRepo Repository.UserRepositoryInterface) Service.UserServiceInterface {
	return &UserService{userRepository: userRepo}
}
