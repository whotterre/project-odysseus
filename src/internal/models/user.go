package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                         uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Email                      string         `gorm:"type:varchar(255);uniqueIndex;not null"`
	PasswordHash               string         `gorm:"not null"`
	AccountPublicKey           string         `gorm:"not null"`
	EncryptedAccountPrivateKey string         `gorm:"type:text;not null"`
	DevicePublicKey            string         `gorm:"not null"`
	
	CreatedAt                  time.Time
	UpdatedAt                  time.Time
	DeletedAt                  gorm.DeletedAt `gorm:"index"`
}