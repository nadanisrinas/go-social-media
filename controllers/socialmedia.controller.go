package controllers

import (
	"gosocialmedia/errs"
	"gosocialmedia/models"
	"gosocialmedia/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SocialMediaController struct {
	socialmediaService services.SocialMediaService
}

func NewSocialMediaService(socialmediaService services.SocialMediaService) *SocialMediaController {
	return &SocialMediaController{socialmediaService: socialmediaService}
}

func (s *SocialMediaController) CreateSocialMedia(ctx *gin.Context) {
	userData, ok := ctx.MustGet("userData").(*models.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	// var requestBody models.CreateSocialMediaRequest
	var requestBody models.SocialMedia

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	createdSocialmedia, err := s.socialmediaService.CreateSocialMedia(userData, &requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusCreated, createdSocialmedia)
}

func (s *SocialMediaController) GetAllSocialMediasByUserSosmed(ctx *gin.Context) {
	userData, ok := ctx.MustGet("userData").(*models.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	socialmedias, err := s.socialmediaService.GetAllSocialMediasByUserSosmed(userData.ID)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, socialmedias)
}

func (s *SocialMediaController) UpdateSocialMedia(ctx *gin.Context) {
	// socialMediaID := ctx.Param("socialMediaID")
	// socialMediaIDUint, err := strconv.ParseUint(socialMediaID, 10, 32)
	// if err != nil {
	// 	errValidation := errs.NewBadRequest("Social Media id should be in unsigned integer")
	// 	ctx.JSON(errValidation.StatusCode(), errValidation)
	// 	return
	// }

	// var reqBody models.UpdateSocialMediaRequest
	var reqBody models.SocialMedia

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		errValidation := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(errValidation.StatusCode(), errValidation)
		return
	}

	updatedSocialMedia, errUpdate := s.socialmediaService.UpdateSocialMedia(&reqBody, &reqBody)
	if errUpdate != nil {
		ctx.JSON(errUpdate.StatusCode(), errUpdate)
		return
	}

	ctx.JSON(http.StatusOK, updatedSocialMedia)
}

func (s *SocialMediaController) DeleteSocialMedia(ctx *gin.Context) {
	socialMediaID := ctx.Param("socialMediaID")
	socialMediaIDUint, err := strconv.ParseUint(socialMediaID, 10, 32)
	if err != nil {
		errValidation := errs.NewBadRequest("Comment id should be in unsigned integer")
		ctx.JSON(errValidation.StatusCode(), errValidation)
		return
	}

	errDelete := s.socialmediaService.DeleteSocialMedia(uint(socialMediaIDUint))
	if errDelete != nil {
		ctx.JSON(errDelete.StatusCode(), errDelete)
		return
	}

	ctx.JSON(http.StatusOK, "success delete social media")
}
