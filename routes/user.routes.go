package routes

import (
	"gosocialmedia/controllers"
	"gosocialmedia/services"

	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController *controllers.UserController
}

func NewUserRouteController(userController *controllers.UserController) *UserRouteController {
	return &UserRouteController{userController}
}

func (oc *UserRouteController) UserRoute(rg *gin.RouterGroup, userService services.UserService) {

	router := rg.Group("users")
	router.POST("/register", oc.userController.Register)
	router.POST("/login", oc.userController.Login)
	router.PUT("", userService.Authentication(), oc.userController.UpdateUser)
	router.DELETE("", userService.Authentication(), oc.userController.DeleteUser)
}
