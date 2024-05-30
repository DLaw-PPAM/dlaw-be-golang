package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LawyerUser struct {
	ID             uuid.UUID      `json:"id" gorm:"primary_key;unique;type:uuid;default:uuid_generate_v4()"`
	ClientID       uuid.UUID      `json:"client_id"`
	PricePerHour   int            `json:"price_per_hour"`
	Rating         int            `json:"rating"`
	ProfilePicture string         `json:"profile_picture"`
	Specialties    []Specialties  `gorm:"many2many:lawyer_specialties;"`
	User           User           `gorm:"foreignKey:ClientID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ClientName     string         `json:"client_name"`
	Reviews        []Review       `gorm:"many2many:lawyer_review;"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
