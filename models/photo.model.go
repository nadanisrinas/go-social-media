package models

import (
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title    string    `json:"title" gorm:"not null"`
	Caption  string    `json:"caption"`
	PhotoURL string    `json:"photo_url" gorm:"not null"`
	UserID   uint      `json:"user_id"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type CreatePhotoRequest struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" binding:"required,url"`
}

func (p *CreatePhotoRequest) ToEntity() *Photo {
	return &Photo{
		Title:    p.Title,
		Caption:  p.Caption,
		PhotoURL: p.PhotoURL,
	}
}

type CreatePhotoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GetAllPhotosResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      UserData  `json:"user"`
}

type UserData struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UpdatePhotoRequest CreatePhotoRequest

func (p *UpdatePhotoRequest) ToEntity() *Photo {
	return &Photo{
		Title:    p.Title,
		Caption:  p.Caption,
		PhotoURL: p.PhotoURL,
	}
}

type UpdatePhotoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeletePhotoResponse struct {
	Message string `json:"message"`
}
