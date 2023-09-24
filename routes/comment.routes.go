package routes

import (
	"gosocialmedia/controllers"
	"gosocialmedia/services"

	"github.com/gin-gonic/gin"
)

type CommentRouteController struct {
	commentController *controllers.CommentController
}

func NewCommentRouteController(commentController *controllers.CommentController) *CommentRouteController {
	return &CommentRouteController{commentController}
}

func (oc *CommentRouteController) CommentRoute(rg *gin.RouterGroup, commentService services.CommentService) {
	router := rg.Group("comment")
	router.GET("", oc.commentController.GetAllComments)
	router.POST("", oc.commentController.CreateComment)
	router.DELETE("/:id", oc.commentController.DeleteComment)
	router.PUT("/:id", oc.commentController.UpdateComment)

}
