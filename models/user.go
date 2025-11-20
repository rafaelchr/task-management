package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	InternalID int64          `json:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicID   uuid.UUID      `json:"public_id" gorm:"unique;column:public_id"`
	Name       string         `json:"name" gorm:"column:name;not null"`
	Email      string         `json:"email" gorm:"unique;column:email;not null"`
	Password   string         `json:"password" gorm:"column:password;not null"`
	Role       string         `json:"role" gorm:"column:role;default:user"`
	CreatedAt  time.Time      `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"column:deleted_at;index"`
}
