package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	InternalID   int64     `json:"internal_id" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicID     uuid.UUID `json:"public_id" db:"public_id" gorm:"column:public_id"`
	CardID       int64     `json:"card_internal_id" db:"card_internal_id" gorm:"column:card_internal_id"`
	CardPublicID uuid.UUID `json:"card_public_id" db:"board_public_id" gorm:"column:card_public_id"`
	UserID       int64     `json:"user_internal_id" db:"user_internal_id" gorm:"column:user_internal_id"`
	UserPublicID uuid.UUID `json:"user_public_id" db:"user_public_id" gorm:"column:user_public_id"`
	Message      string    `json:"message" db:"message" gorm:"column:message"`
	CreatedAt    time.Time `json:"created_at" db:"created_at" gorm:"column:created_at"`
}
