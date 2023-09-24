package routes

import (
	"gosocialmedia/controllers"
	"gosocialmedia/services"

	"github.com/gin-gonic/gin"
)

type PhotoRouteController struct {
	photoController *controllers.PhotoController
	userController  *controllers.UserController
}

func NewPhotoRouteController(PhotoController *controllers.PhotoController, userController *controllers.UserController) *PhotoRouteController {
	return &PhotoRouteController{PhotoController, userController}
}

func (oc *PhotoRouteController) PhotoRoute(rg *gin.RouterGroup, PhotoService services.PhotoService, userService services.UserService) {

	router := rg.Group("photo")
	router.POST("/photos", userService.Authentication(), oc.photoController.CreatePhoto)
	router.GET("/photos", userService.Authentication(), oc.photoController.GetAllPhotos)
	router.PUT("/photos/:photoID", userService.Authentication(), oc.userController.PhotosAuthorization(), oc.photoController.UpdatePhoto)
	router.DELETE("/photos/:photoID", userService.Authentication(), oc.userController.PhotosAuthorization(), oc.photoController.DeletePhoto)

}
