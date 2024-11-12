package repositories

import (
	"errors"
	"log"
	"server/dtos/user"
	"server/interface/Repository"
	"server/models"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (u *UserRepository) GetCountUser() (int64, error) {
	var count int64
	if err := u.DB.Table("user").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (u *UserRepository) GetUserById(userId uint) (*models.User, error) {
	var user *models.User
	if err := u.DB.Preload("Roles").First(&user, userId).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) GetAllUsers(fullName, role string) ([]*models.User, error) {
	var user []*models.User
	query := u.DB.Model(&models.User{}).Preload("Roles")
	if fullName != "" {
		query = query.Where("full_name LIKE ?", "%"+fullName+"%")
	}

	if role != "" {
		query = query.Joins("JOIN user_roles ON user_roles.user_id = \"user\".id").
			Joins("JOIN role ON role.id = user_roles.role_id").
			Where("role.role_name = ?", role)
	}

	if err := query.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) DeleteUser(userId uint) error {
	result := u.DB.Table("user").Where("id = ?", userId).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
func (u *UserRepository) UpdateUser(userId uint, dto user.UpdateUserDto) (*models.User, error) {
	var existingUser models.User

	// Tìm người dùng
	if err := u.DB.First(&existingUser, userId).Error; err != nil {
		log.Printf("User not found: %v", err)
		return nil, err
	}

	// Tạo map các trường để cập nhật
	updates := map[string]interface{}{}
	if dto.FullName != "" {
		updates["full_name"] = dto.FullName
	}
	if dto.Email != "" {
		updates["email"] = dto.Email
	}
	if dto.Phone != "" {
		updates["phone"] = dto.Phone
	}
	if dto.ImageUrl != "" {
		updates["image_url"] = dto.ImageUrl
	}

	// Cập nhật các trường vào database
	if len(updates) > 0 {
		if err := u.DB.Model(&existingUser).Updates(updates).Error; err != nil {
			log.Printf("Error updating user: %v", err)
			return nil, err
		}
	}

	// Cập nhật roles nếu có
	if len(dto.Roles) > 0 {
		var roles []models.Role
		if err := u.DB.Table("role").Where("id IN ?", dto.Roles).Find(&roles).Error; err != nil {
			log.Printf("Error retrieving roles: %v", err)
			return nil, err
		}
		if err := u.DB.Model(&existingUser).Association("Roles").Replace(&roles); err != nil {
			log.Printf("Error updating roles for user: %v", err)
			return nil, err
		}
	}

	// Tải lại thông tin người dùng với roles
	if err := u.DB.Preload("Roles").First(&existingUser, userId).Error; err != nil {
		log.Printf("Error retrieving updated user: %v", err)
		return nil, err
	}

	return &existingUser, nil
}

//func (u *UserRepository) UpdateUser(userId uint, dto user.UpdateUserDto) (*models.User, error) {
//	var existingUser models.User
//
//	if err := u.DB.First(&existingUser, userId).Error; err != nil {
//		log.Printf("User not found: %v", err)
//		return nil, err
//	}
//
//	updates := map[string]interface{}{
//		"full_name": dto.FullName,
//		"email":     dto.Email,
//		"phone":     dto.Phone,
//	}
//
//	if err := u.DB.Model(&existingUser).Updates(updates).Error; err != nil {
//		log.Printf("Error updating user: %v", err)
//		return nil, err
//	}
//
//	if len(dto.Roles) > 0 {
//		var roles []models.Role
//		if err := u.DB.Table("role").Where("id IN ?", dto.Roles).Find(&roles).Error; err != nil {
//			log.Printf("Error retrieving roles: %v", err)
//			return nil, err
//		}
//		if err := u.DB.Model(&existingUser).Association("Roles").Replace(&roles); err != nil {
//			log.Printf("Error updating roles for user: %v", err)
//			return nil, err
//		}
//	}
//
//	if err := u.DB.Preload("Roles").First(&existingUser, userId).Error; err != nil {
//		log.Printf("Error retrieving updated user: %v", err)
//		return nil, err
//	}
//
//	return &existingUser, nil
//}

func (u *UserRepository) FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := u.DB.Table("user").Where("username = ?", username).Preload("Roles").
		Preload("Roles.Permissions").First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) Register(userRegisterDto *user.UserRegister) (*models.User, error) {
	var roles []models.Role
	if len(userRegisterDto.Roles) > 0 {
		if err := u.DB.Where("id IN ?", userRegisterDto.Roles).Find(&roles).Error; err != nil {
			log.Printf("Roles not found: %v", err)
			return nil, err
		}
	}

	newUser := models.User{
		Username: userRegisterDto.Username,
		Password: userRegisterDto.Password,
		FullName: userRegisterDto.FullName,
		Email:    userRegisterDto.Email,
		Phone:    userRegisterDto.Phone,
		Roles:    roles,
	}

	if err := u.DB.Create(&newUser).Error; err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}

	return &newUser, nil
}

func NewUserRepository(db *gorm.DB) Repository.UserRepositoryInterface {
	return &UserRepository{
		DB: db,
	}
}
