package controllers

import (
	"gosocialmedia/errs"
	"gosocialmedia/models"
	"gosocialmedia/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService  services.UserService
	photoService services.PhotoService
}

func NewUserController(userService services.UserService, photoService services.PhotoService) *UserController {
	return &UserController{userService: userService, photoService: photoService}
}

func (u *UserController) Register(ctx *gin.Context) {
	var requestBody models.User

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	registeredUser, err := u.userService.Register(&requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusCreated, registeredUser)
}

func (u *UserController) Login(ctx *gin.Context) {
	var requestBody models.LoginRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	token, err := u.userService.Login(&requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, token)
}

func (u *UserController) UpdateUser(ctx *gin.Context) {
	var requestBody models.User
	userData, ok := ctx.MustGet("userData").(*models.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	updatedUser, err := u.userService.UpdateUser(userData, &requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, updatedUser)
}

func (u *UserController) DeleteUser(ctx *gin.Context) {
	userData, ok := ctx.MustGet("userData").(*models.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	err := u.userService.DeleteUser(userData.ID)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, "berhasil delete user")
}

func (u *UserController) PhotosAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, ok := ctx.MustGet("userData").(*models.User)
		if !ok {
			newError := errs.NewBadRequest("Failed to get user data")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		photoID := ctx.Param("photoID")
		photoIDUint, err := strconv.ParseUint(photoID, 10, 32)
		if err != nil {
			newError := errs.NewBadRequest("Photo id should be an unsigned integer")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		photo, err2 := u.photoService.GetPhotoByID(uint(photoIDUint))
		if err2 != nil {
			ctx.AbortWithStatusJSON(err2.StatusCode(), err2)
			return
		}

		if photo.UserID != userData.ID {
			newError := errs.NewUnauthorized("You're not authorized to modify this photo")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		ctx.Next()
	}
}
