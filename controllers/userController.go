package controllers

import (
	"net/http"
	"os"
	"server/dtos/user"
	"server/interface/Service"
	"server/models"
	"server/utils"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

type UserController struct {
	userService Service.UserServiceInterface
}

func NewUserController(userService Service.UserServiceInterface) *UserController {
	return &UserController{userService: userService}
}

func (u *UserController) GetProfile(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}
	claims := user.(jwt.MapClaims)
	c.JSON(http.StatusOK, gin.H{"user": claims})
}

func (u *UserController) Register(c *gin.Context) {
	var userRegisterDto user.UserRegister
	if err := c.ShouldBind(&userRegisterDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	if err := userRegisterDto.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	hassedPassword, err := bcrypt.GenerateFromPassword([]byte(userRegisterDto.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Message: "HasPassword failure",
			Status:  http.StatusBadRequest,
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	userRegisterDto.Password = string(hassedPassword)

	user, err := u.userService.Register(&userRegisterDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "User registration failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "User registered successfully",
		Data:    user,
		Error:   "",
	})
}

func (u *UserController) Login(c *gin.Context) {
	var userLoginDto user.UserLoginDto
	if err := c.ShouldBind(&userLoginDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	user, err := u.userService.FindUserByUsername(userLoginDto.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, &utils.Response{
			Status:  http.StatusNotFound,
			Message: "User not found",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLoginDto.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, &utils.Response{
			Status:  http.StatusUnauthorized,
			Message: "Invalid credentials",
			Data:    nil,
			Error:   "Password mismatch",
		})
		return
	}

	token, err := generateJWT(uint(user.ID), user.Roles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "Token generation error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	type LoginResponse struct {
		Token string       `json:"token"`
		User  *models.User `json:"user"`
	}

	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "User logged in",
		Data: LoginResponse{
			User:  user,
			Token: token,
		},
		Error: "",
	})
}

func generateJWT(userID uint, roles []models.Role) (string, error) {
	var roleNames []string
	var permissionNames []string
	for _, role := range roles {
		roleNames = append(roleNames, role.RoleName)
		for _, permission := range role.Permissions {
			permissionNames = append(permissionNames, permission.PermissionName)
		}
	}

	claims := jwt.MapClaims{
		"id":          userID,
		"roles":       roleNames,
		"permissions": permissionNames,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func (u *UserController) UpdateUser(c *gin.Context) {
	var userUpdateDto user.UpdateUserDto
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)

	if err := c.ShouldBind(&userUpdateDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	user, err := u.userService.UpdateUser(uint(userId), userUpdateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "User update error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "User updated successfully",
		Data:    user,
		Error:   "",
	})
}

func (u *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("userId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	if err := u.userService.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "User delete error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "User deleted successfully",
		Data:    nil,
		Error:   "",
	})
	return
}

func (u *UserController) GetAllUsers(c *gin.Context) {
	fullName := c.Query("full_name")
	role := c.Query("role")
	users, err := u.userService.GetAllUsers(fullName, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "User get all error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "User get all",
		Data:    users,
		Error:   "",
	})
	return
}

func (u *UserController) GetUserById(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	user, err := u.userService.GetUserById(uint(uint(userId)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "User get error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "User get successfully",
		Data:    user,
		Error:   "",
	})
}
