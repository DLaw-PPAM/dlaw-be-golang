package dto

import (
	"example.com/m/v2/model/entity"
	"github.com/google/uuid"
)

type AddLawyerRequestDTO struct {
	ClientID     uuid.UUID            `json:"client_id"`
	PricePerHour int                  `json:"price_per_hour" validate:"required"`
	Rating       int                  `json:"rating" validate:"required"`
	Specialties  []entity.Specialties `json:"specialties"`
}

type AddLawyerResponseDTO struct {
	Message      string               `json:"message"`
	ID           uuid.UUID            `json:"id"`
	ClientID     uuid.UUID            `json:"client_id"`
	PricePerHour int                  `json:"price_per_hour"`
	Rating       int                  `json:"rating"`
	Specialties  []entity.Specialties `json:"specialties"`
	User         entity.User          `json:"user"`
}
