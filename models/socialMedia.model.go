package models

import (
	"time"

	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model
	Name           string `json:"name" gorm:"not null"`
	SocialMediaURL string `json:"social_media_url" gorm:"not null"`
	UserID         uint   `json:"user_id"`
}
type CreateSocialMediaRequest struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaURL string `json:"social_media_url" binding:"required,url"`
}

func (s *CreateSocialMediaRequest) ToEntity() *SocialMedia {
	return &SocialMedia{
		Name:           s.Name,
		SocialMediaURL: s.SocialMediaURL,
	}
}

type CreateSocialMediaResponse struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type GetAllSocialMediasResponse struct {
	SocialMedias []SocialMediaData `json:"social_medias"`
}

type UserDataSocialMedia struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type UpdateSocialMediaRequest struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaURL string `json:"social_media_url" binding:"required,url"`
}

func (s *UpdateSocialMediaRequest) ToEntity() *SocialMedia {
	return &SocialMedia{
		Name:           s.Name,
		SocialMediaURL: s.SocialMediaURL,
	}
}

type UpdateSocialMediaResponse struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type SocialMediaData struct {
	ID             uint                `json:"id"`
	Name           string              `json:"name"`
	SocialMediaURL string              `json:"social_media_url"`
	UserID         uint                `json:"user_id"`
	CreatedAt      time.Time           `json:"created_at"`
	UpdatedAt      time.Time           `json:"updated_at"`
	User           UserDataSocialMedia `json:"user"`
}

type DeleteSocialMediaResponse struct {
	Message string `json:"message"`
}
