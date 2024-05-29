package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DetailLawyer struct {
	ID           uuid.UUID      `json:"id" gorm:"primary_key;unique;type:uuid;default:uuid_generate_v4()"`
	ClientID     uuid.UUID      `json:"client_id"`
	PricePerHour int            `json:"price_per_hour"`
	Rating       int            `json:"rating"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
