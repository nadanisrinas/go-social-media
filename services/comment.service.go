package services

import (
	"gosocialmedia/models"
)

type CommentService interface {
	GetAllComments() (*models.Comment, error)
	CreateComments(message string, photoId int64) (*models.Comment, error)
	FindComment(commentId int64) (models.Comment, error)
	DeleteComment(idComment int64) (*models.Comment, error)
	UpdateComment(message string, commentId int64) (*models.Comment, error)
}
