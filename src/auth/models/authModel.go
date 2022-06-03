package models

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/yamess/auth/database"
	"gorm.io/gorm"
	"io"
)

type UserCredentials struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required"`
}

type Token struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type LogOut struct {
	Token string `json:"token"`
}

type BlackListToken struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Token     string `json:"token" validate:"required" gorm:"unique"`
	CreatedAt NullTime
	CreatedBy uuid.UUID
}

func (b *BlackListToken) CreateRecord(userId uuid.UUID) *gorm.DB {
	b.CreatedBy = userId
	result := database.MyDB.Conn.Where(BlackListToken{Token: b.Token}).FirstOrCreate(&b)
	return result
}

// FromJSON and ToJSON encoder decoder
func (c *UserCredentials) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(&c)
}

func (c *UserCredentials) Validate() error {
	validation := validator.New()
	return validation.Struct(c)
}
