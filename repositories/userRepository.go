package repositories

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"server/dtos/user"
	"server/interface/Repository"
	"server/models"
	"time"
)

type UserRepository struct {
	DB *gorm.DB
}

func (u *UserRepository) GetAllUsers(fullName string) ([]*models.User, error) {
	var users []*models.User
	query := u.DB.Model(&models.User{})
	if fullName != "" {
		query = query.Where("full_name LIKE ?", "%"+fullName+"%")
	}
	if err := u.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepository) DeleteUser(userId int) error {
	result := u.DB.Table("users").Where("id = ?", userId).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (u *UserRepository) UpdateUser(userId int, dto user.UserUpdateDto) (*models.User, error) {
	var existingUser models.User
	if err := u.DB.First(&existingUser, userId).Error; err != nil {
		log.Printf("User not found: %v", err)
		return nil, err
	}

	updates := map[string]interface{}{
		"full_name": dto.FullName,
		"email":     dto.Email,
		"phone":     dto.Phone,
		"role":      dto.Role,
	}

	if err := u.DB.Model(&existingUser).Updates(updates).Error; err != nil {
		log.Printf("Error updating user: %v", err)
		return nil, err
	}

	if err := u.DB.First(&existingUser, userId).Error; err != nil {
		log.Printf("Error retrieving updated user: %v", err)
		return nil, err
	}
	return &existingUser, nil
}

func (u *UserRepository) FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := u.DB.Table("users").Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) Register(userRegisterDto *user.UserRegister) (*models.User, error) {
	if err := u.DB.Table("users").Create(userRegisterDto).Error; err != nil {
		return nil, err
	}

	m := &models.User{
		Username: userRegisterDto.Username,
		Email:    userRegisterDto.Email,
		FullName: userRegisterDto.FullName,
		Phone:    userRegisterDto.Phone,
		Role:     models.Role(userRegisterDto.Role),
	}

	return m, nil
}

func NewUserRepository(db *gorm.DB) Repository.UserRepositoryInterface {
	return &UserRepository{
		DB: db,
	}
}
