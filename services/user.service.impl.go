package services

import (
	"fmt"
	"gosocialmedia/errs"
	"gosocialmedia/models"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &UserServiceImpl{db}
}

func (u *UserServiceImpl) Register(user *models.User) (*models.User, errs.MessageErr) {
	if err := u.db.Create(user).Error; err != nil {
		log.Println(err.Error())
		return nil, errs.NewInternalServerError("Failed to register new user")
	}

	return user, nil
}

func (u *UserServiceImpl) GetUserByEmail(email string) (*models.User, errs.MessageErr) {
	var user models.User

	if err := u.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("User with email %s is not found", email))
	}

	return &user, nil
}

func (u *UserServiceImpl) GetUserByID(id uint) (*models.User, errs.MessageErr) {
	var user models.User

	if err := u.db.First(&user, id).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("User with id %v is not found", id))
	}

	return &user, nil
}

func (u *UserServiceImpl) GetSocialMediaByID(id uint) (*models.SocialMedia, errs.MessageErr) {
	var socialmedia models.SocialMedia
	if err := u.db.First(&socialmedia, id).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("Social Media with id %d is not found", id))
	}

	return &socialmedia, nil
}

func (u *UserServiceImpl) UpdateUser(oldUser *models.User, newUser *models.User) (*models.User, errs.MessageErr) {
	if err := u.db.Model(oldUser).Updates(newUser).Error; err != nil {
		return nil, errs.NewBadRequest(fmt.Sprintf("Failed to update user with id %v", oldUser.ID))
	}

	return oldUser, nil
}

func (u *UserServiceImpl) DeleteUser(id uint) errs.MessageErr {
	if err := u.db.Delete(&models.User{}, id).Error; err != nil {
		return errs.NewInternalServerError(fmt.Sprintf("Failed to delete user with id %d", id))
	}

	return nil
}

func (u *UserServiceImpl) Login(payload *models.LoginRequest) (*models.LoginResponse, errs.MessageErr) {
	user, err := u.GetUserByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	if err := user.ComparePassword(payload.Password); err != nil {
		return nil, err
	}

	token, err2 := user.CreateToken()
	if err2 != nil {
		return nil, err2
	}

	response := &models.LoginResponse{
		Token: token,
	}

	return response, nil
}

func (u *UserServiceImpl) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := ctx.GetHeader("Authorization")

		var user models.User

		if err := user.ValidateToken(bearerToken); err != nil {
			ctx.AbortWithStatusJSON(err.StatusCode(), err)
			return
		}

		result, err := u.GetUserByID(user.ID)
		if err != nil {
			ctx.AbortWithStatusJSON(err.StatusCode(), err)
			return
		}

		ctx.Set("userData", result)
		ctx.Next()
	}
}

func (u *UserServiceImpl) SocialmediasAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, ok := ctx.MustGet("userData").(*models.User)
		if !ok {
			newError := errs.NewBadRequest("Failed to get user data")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		socialmediaID := ctx.Param("socialMediaID")
		socialmediaIDUint, err := strconv.ParseUint(socialmediaID, 10, 32)
		if err != nil {
			newError := errs.NewBadRequest("Social Media id should be an unsigned integer")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		socialmedia, err2 := u.GetSocialMediaByID(uint(socialmediaIDUint))
		if err2 != nil {
			ctx.AbortWithStatusJSON(err2.StatusCode(), err2)
			return
		}

		if socialmedia.UserID != userData.ID {
			newError := errs.NewUnauthorized("You're not authorized to modify this Social Media")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		ctx.Next()
	}
}
