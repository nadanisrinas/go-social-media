package services

import (
	"gosocialmedia/errs"
	"gosocialmedia/models"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Register(user *models.User) (*models.User, errs.MessageErr)
	Login(payload *models.LoginRequest) (*models.LoginResponse, errs.MessageErr)
	UpdateUser(oldUser *models.User, newUser *models.User) (*models.User, errs.MessageErr)
	DeleteUser(id uint) errs.MessageErr
	// GetUserByEmail(email string) (*models.User, errs.MessageErr)
	GetUserByID(id uint) (*models.User, errs.MessageErr)
	Authentication() gin.HandlerFunc
	SocialmediasAuthorization() gin.HandlerFunc
}
