package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID    uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name  string    `json:"name" gorm:"type:varchar(100);not null"`
	Email string    `json:"email" gorm:"type:varchar(100);unique;not null"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New()
	return nil
}
