package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	ID          uuid.UUID      `json:"id" gorm:"primary_key;unique;type:uuid;default:uuid_generate_v4()"`
	Rating      int            `json:"rating"`
	Description string         `json:"description"`
	LawyerID    uuid.UUID      `json:"lawyer_id"`
	ClientID    uuid.UUID      `json:"client_id"`
	ClientName  string         `json:"client_name"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
