package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Case struct {
	ID          uuid.UUID      `json:"id" gorm:"primary_key;unique;type:uuid;default:uuid_generate_v4()"`
	Subject     string         `json:"subject"`
	Media       string         `json:"media"`
	Notes       string         `json:"notes"`
	Status      string         `json:"status"`
	Hour        int            `json:"hour"`
	AdditionFee int            `json:"addition_fee"`
	Lawyer      LawyerUser     `json:"lawyer_id" `
	ClientID    uuid.UUID      `json:"client_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
