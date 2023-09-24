package services

import (
	"fmt"
	"gosocialmedia/models"
	"log"

	"gorm.io/gorm"
)

type CommentServiceImpl struct {
	db *gorm.DB
}

func NewCommentService(db *gorm.DB) CommentService {
	return &CommentServiceImpl{db}
}

func (csi *CommentServiceImpl) GetAllComments() (*models.Comment, error) {
	comments := models.Comment{}
	result := csi.db.Find(&comments)
	if result.Error != nil {
		log.Fatal("Error get items data", result.Error)
	}

	return &comments, result.Error
}

func (csi *CommentServiceImpl) CreateComments(message string, photoId int64) (*models.Comment, error) {
	payload := models.Comment{
		Message: message,
	}
	err := csi.db.Create(&payload).Error
	if err != nil {
		log.Fatal("Error create items data", err)
	}

	fmt.Println("New item data:", payload)
	return &payload, err
}

func (csi *CommentServiceImpl) FindComment(commentId int64) (models.Comment, error) {
	// fmt.Println("itemCode", itemCode)
	comments := models.Comment{}
	// itemCodeUUID, errUUID := uuid.FromString(itemCode)
	errFindComment := csi.db.Where("id = ?", commentId).Find(&comments).Error
	if errFindComment != nil {
		log.Fatal("Error can't find item", errFindComment)
	}

	return comments, errFindComment
}

func (csi *CommentServiceImpl) DeleteComment(idComment int64) (*models.Comment, error) {
	orders := &models.Comment{}
	result := csi.db.First(orders, idComment)
	if result.Error != nil {
		log.Fatal("Error get orders data", result.Error)
	}
	result = csi.db.Delete(orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, result.Error
}

func (csi *CommentServiceImpl) UpdateComment(message string, commentId int64) (*models.Comment, error) {
	comments := &models.Comment{}
	const layout = "01-02-2006"
	result := csi.db.First(comments, commentId)
	if result.Error != nil {
		log.Fatal("Error get comments data", result.Error)
	}

	payload := models.Comment{
		Message: message,
	}
	result = csi.db.Save(&payload)
	if result.Error != nil {
		return nil, result.Error
	}

	return comments, result.Error
}
