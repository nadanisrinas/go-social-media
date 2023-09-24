package services

import (
	"fmt"
	"gosocialmedia/errs"
	"gosocialmedia/models"
	"log"

	"gorm.io/gorm"
)

type socialmediaImpl struct {
	db *gorm.DB
}

func NewSocialmediaService(db *gorm.DB) SocialMediaService {
	return &socialmediaImpl{db: db}
}

func (s *socialmediaImpl) CreateSocialMedia(user *models.User, socialmedia *models.SocialMedia) (*models.SocialMedia, errs.MessageErr) {
	if err := s.db.Model(user).Association("SocialMedias").Append(socialmedia); err != nil {
		log.Println(err.Error())
		return nil, errs.NewInternalServerError("Failed to create new Social Media")
	}
	return socialmedia, nil
}

func (s *socialmediaImpl) GetAllSocialMediasByUserSosmed(userID uint) ([]models.SocialMedia, errs.MessageErr) {
	var socialmedias []models.SocialMedia
	if err := s.db.Find(&socialmedias, "user_id = ?", userID).Error; err != nil {
		log.Println(err.Error())
		return nil, errs.NewInternalServerError("Failed to get all social media")
	}

	return socialmedias, nil
}

func (s *socialmediaImpl) GetSocialMediaByID(id uint) (*models.SocialMedia, errs.MessageErr) {
	var socialmedia models.SocialMedia
	if err := s.db.First(&socialmedia, id).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("Social Media with id %d is not found", id))
	}

	return &socialmedia, nil
}

func (s *socialmediaImpl) UpdateSocialMedia(oldSocialMedia *models.SocialMedia, newSocialMedia *models.SocialMedia) (*models.SocialMedia, errs.MessageErr) {
	if err := s.db.Model(oldSocialMedia).Updates(newSocialMedia).Error; err != nil {
		log.Println(err.Error())
		return nil, errs.NewInternalServerError(fmt.Sprintf("Failed to update social media with id %d", oldSocialMedia.ID))
	}

	return oldSocialMedia, nil
}

func (s *socialmediaImpl) DeleteSocialMedia(id uint) errs.MessageErr {
	if err := s.db.Delete(&models.SocialMedia{}, id).Error; err != nil {
		log.Println("Error:", err.Error())
		return errs.NewInternalServerError(fmt.Sprintf("Failed to delete Social Media with id %d", id))
	}

	return nil
}
