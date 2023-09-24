package models

import (
	"gosocialmedia/errs"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var JWT_SECRET = os.Getenv("JWT_SECRET")

type User struct {
	gorm.Model
	Username     string        `gorm:"unique;not null" binding:"required"`
	Email        string        `gorm:"unique;not null" binding:"email,required"`
	Password     string        `gorm:"not null" binding:"required,min=6"`
	Age          uint          `gorm:"not null" binding:"required,min=8"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"email,required"`
	Password string `json:"password" binding:"required,min=6"`
	Age      uint   `json:"age" binding:"required,min=8"`
}

func (r *User) ToEntity() User {
	return User{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
		Age:      r.Age,
	}
}

type RegisterResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"email,required"`
	Age      uint   `json:"age" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"email,required"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	Token string `json:"token" binding:"jwt"`
}

type UpdateUserRequest struct {
	Email    string `json:"email" binding:"email,required"`
	Username string `json:"username" binding:"required"`
}

type UpdateUserResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email" binding:"email,required"`
	Username  string    `json:"username" binding:"required"`
	Age       uint      `json:"age" binding:"required,min=8"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteUserRequest = LoginResponse

type DeleteUserResponse struct {
	Message string `json:"message"`
}

func (u *User) HashPassword() errs.MessageErr {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return errs.NewInternalServerError("Failed to hash password")
	}

	u.Password = string(hashedPassword)

	return nil
}

func (u *User) ComparePassword(password string) errs.MessageErr {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return errs.NewBadRequest("Password is not valid!")
	}

	return nil
}

func (u *User) CreateToken() (string, errs.MessageErr) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": u.ID,
			"exp":    time.Now().Add(1 * time.Hour).Unix(),
		})

	signedString, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		log.Println("Error:", err.Error())
		return "", errs.NewInternalServerError("Failed to sign jwt token")
	}

	return signedString, nil
}

func (u *User) ValidateToken(bearerToken string) errs.MessageErr {
	if isBearer := strings.HasPrefix(bearerToken, "Bearer"); !isBearer {
		return errs.NewUnauthenticated("Token type should be Bearer")
	}

	splitToken := strings.Fields(bearerToken)
	if len(splitToken) != 2 {
		return errs.NewUnauthenticated("Token is not valid")
	}

	tokenString := splitToken[1]
	token, err := u.ParseToken(tokenString)
	if err != nil {
		return err
	}

	var mapClaims jwt.MapClaims

	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return errs.NewUnauthenticated("Token is not valid")
	} else {
		mapClaims = claims
	}

	return u.bindTokenToUserEntity(mapClaims)
}

func (u *User) ParseToken(tokenString string) (*jwt.Token, errs.MessageErr) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.NewUnauthenticated("Token method is not valid")
		}
		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		return nil, errs.NewUnauthenticated("Token is not valid")
	}

	return token, nil
}

func (u *User) bindTokenToUserEntity(claim jwt.MapClaims) errs.MessageErr {
	if id, ok := claim["userId"].(float64); !ok {
		return errs.NewUnauthenticated("Token doesn't contains userId")
	} else {
		u.ID = uint(id)
	}

	return nil
}
