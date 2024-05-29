package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID             uuid.UUID      `json:"id" gorm:"primary_key;unique;type:uuid;default:uuid_generate_v4()"`
	Username       string         `json:"username"`
	ProfilePicture string         `json:"profile_picture"`
	Email          string         `json:"email"`
	Password       string         `json:"password"`
	FullName       string         `json:"full_name"`
	BirthDate      time.Time      `json:"birth_date"`
	PhoneNumber    string         `json:"phone_number"`
	Address        string         `json:"address"`
	Bio            string         `json:"bio"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
