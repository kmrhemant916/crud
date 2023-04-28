package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `gorm:"type:char(36);primary_key;"`
	Email  string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	CreatedAt  time.Time `gorm:"default:0" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:0" json:"updated_at"`
}