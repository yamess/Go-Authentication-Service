package models

import (
	"github.com/google/uuid"
	"time"
)

type Base struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt NullTime  `json:"updated_at" gorm:"default:null"`
	CreatedBy uuid.UUID `json:"created_by" example:"2d21b581-715e-4b3a-a778-ae1db0aa445a"`
	UpdatedBy uuid.UUID `json:"updated_by" gorm:"default:null" example:"2d21b581-715e-4b3a-a778-ae1db0aa445a"`
}
