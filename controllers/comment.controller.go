package controllers

import (
	"fmt"
	"gosocialmedia/models"
	"gosocialmedia/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService services.CommentService
}

func NewCommentController(commentService services.CommentService) *CommentController {
	return &CommentController{commentService}
}

func (cs *CommentController) GetAllComments(ctx *gin.Context) {
	comments, err := cs.commentService.GetAllComments()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"comments": comments}})
}

func (cs *CommentController) DeleteComment(ctx *gin.Context) {
	commentID, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	response, err := cs.commentService.FindComment(commentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"orders": response}})
}

func (cs *CommentController) UpdateComment(ctx *gin.Context) {
	commentID, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	var comment models.Comment
	if err := ctx.BindJSON(&comment); err != nil {
		fmt.Println("! nil")

		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	response, err := cs.commentService.UpdateComment(comment.Message, commentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"orders": response}})
}

func (cs *CommentController) CreateComment(ctx *gin.Context) {
	var OrderRequestBody models.Comment
	// var errItems error
	if err := ctx.BindJSON(&OrderRequestBody); err != nil {
		fmt.Println("! nil")

		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

}
