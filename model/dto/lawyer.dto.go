package dto

import (
	"example.com/m/v2/model/entity"
	"github.com/google/uuid"
)

type AddLawyerRequestDTO struct {
	ClientID       uuid.UUID            `json:"client_id"`
	PricePerHour   int                  `json:"price_per_hour" validate:"required"`
	Rating         int                  `json:"rating" validate:"required"`
	Specialties    []entity.Specialties `json:"specialties"`
	ProfilePicture string               `json:"profile_picture"`
}

type AddLawyerResponseDTO struct {
	Message        string               `json:"message"`
	ID             uuid.UUID            `json:"id"`
	ClientID       uuid.UUID            `json:"client_id"`
	ClientName     string               `json:"client_name"`
	PricePerHour   int                  `json:"price_per_hour"`
	ProfilePicture string               `json:"profile_picture"`
	Rating         int                  `json:"rating"`
	Specialties    []entity.Specialties `json:"specialties"`
	User           entity.User          `json:"user"`
}

type SearchLawyerRequestDTO struct {
	Name string `json:"name"`
}

type SearchLawyerResponseDTO struct {
	Lawyers []entity.LawyerUser `json:"lawyers"`
}
