package services

import (
	"gosocialmedia/errs"
	"gosocialmedia/models"
)

type PhotoService interface {
	CreatePhoto(user *models.User, photo *models.Photo) (*models.Photo, errs.MessageErr)
	GetAllPhotos() ([]models.Photo, errs.MessageErr)
	UpdatePhoto(oldPhoto *models.Photo, newPhoto *models.Photo) (*models.Photo, errs.MessageErr)
	DeletePhoto(id uint) errs.MessageErr
	GetPhotoByID(id uint) (*models.Photo, errs.MessageErr)
}
