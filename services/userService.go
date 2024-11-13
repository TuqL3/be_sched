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

func (u *UserService) GetCountUser() (int64, error) {
	return u.userRepository.GetCountUser()
}

func (u *UserService) GetUserById(userId uint) (*models.User, error) {
	return u.userRepository.GetUserById(userId)
}

func (u *UserService) GetAllUsers(fullName, role string) ([]*models.User, error) {
	return u.userRepository.GetAllUsers(fullName, role)
}

func (u *UserService) DeleteUser(userId uint) error {
	return u.userRepository.DeleteUser(userId)
}

func (u *UserService) UpdateUser(userId uint, dto user.UpdateUserDto, imageUrl string) (*models.User, error) {
	return u.userRepository.UpdateUser(userId, dto, imageUrl)
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
