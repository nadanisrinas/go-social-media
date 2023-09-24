package routes

import (
	"gosocialmedia/controllers"
	"gosocialmedia/services"

	"github.com/gin-gonic/gin"
)

type SocialMediaRouteController struct {
	socialMediaController *controllers.SocialMediaController
	userController        *controllers.UserController
}

func NewsocialMediaRouteController(socialMediaController *controllers.SocialMediaController, userController *controllers.UserController) *SocialMediaRouteController {
	return &SocialMediaRouteController{socialMediaController, userController}
}

func (oc *SocialMediaRouteController) SocialMediaRoute(rg *gin.RouterGroup, socialMediaService services.SocialMediaService, userService services.UserService) {

	router := rg.Group("socialMedia")
	router.POST("", userService.Authentication(), oc.socialMediaController.CreateSocialMedia)
	router.GET("", userService.Authentication(), oc.socialMediaController.GetAllSocialMediasByUserSosmed)
	router.PUT("/:socialMediaID", userService.Authentication(), userService.SocialmediasAuthorization(), oc.socialMediaController.UpdateSocialMedia)
	router.DELETE("/:socialMediaID", userService.Authentication(), userService.SocialmediasAuthorization(), oc.socialMediaController.DeleteSocialMedia)

}
