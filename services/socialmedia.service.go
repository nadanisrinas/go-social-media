package services

import (
	"gosocialmedia/errs"
	"gosocialmedia/models"
)

type SocialMediaService interface {
	CreateSocialMedia(user *models.User, socialmedia *models.SocialMedia) (*models.SocialMedia, errs.MessageErr)
	GetAllSocialMediasByUserSosmed(userID uint) ([]models.SocialMedia, errs.MessageErr)
	GetSocialMediaByID(id uint) (*models.SocialMedia, errs.MessageErr)
	UpdateSocialMedia(oldSocialMedia *models.SocialMedia, newSocialMedia *models.SocialMedia) (*models.SocialMedia, errs.MessageErr)
	DeleteSocialMedia(id uint) errs.MessageErr
}
