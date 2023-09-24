package controllers

import (
	"gosocialmedia/errs"
	"gosocialmedia/models"
	"gosocialmedia/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PhotoController struct {
	photoService services.PhotoService
}

func NewPhotoService(photoService services.PhotoService) *PhotoController {
	return &PhotoController{photoService: photoService}
}

func (p *PhotoController) CreatePhoto(ctx *gin.Context) {
	userData, ok := ctx.MustGet("userData").(*models.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}
	var requestBody models.Photo

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	createdPhoto, err := p.photoService.CreatePhoto(userData, &requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusCreated, createdPhoto)
}

func (p *PhotoController) GetAllPhotos(ctx *gin.Context) {
	photos, err := p.photoService.GetAllPhotos()
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, photos)
}

func (p *PhotoController) UpdatePhoto(ctx *gin.Context) {
	var requestBody models.Photo
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	// photoID := ctx.Param("photoID")
	// photoIDUint, err := strconv.ParseUint(photoID, 10, 32)
	// if err != nil {
	// 	newError := errs.NewBadRequest("Photo id should be an unsigned integer")
	// 	ctx.JSON(newError.StatusCode(), newError)
	// 	return
	// }

	updatedPhoto, err2 := p.photoService.UpdatePhoto(&requestBody, &requestBody)
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	ctx.JSON(http.StatusOK, updatedPhoto)
}

func (p *PhotoController) DeletePhoto(ctx *gin.Context) {
	photoID := ctx.Param("photoID")
	photoIDUint, err := strconv.ParseUint(photoID, 10, 32)
	if err != nil {
		newError := errs.NewBadRequest("Photo id should be an unsigned integer")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	err2 := p.photoService.DeletePhoto(uint(photoIDUint))
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	ctx.JSON(http.StatusOK, "success delete photo")
}
