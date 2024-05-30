package dto

import (
	"github.com/google/uuid"
)

type AddCaseRequestDTO struct {
	Subject     string    `json:"subject" validate:"required"`
	Media       string    `json:"media" validate:"required"`
	Notes       string    `json:"notes" validate:"required"`
	Status      string    `json:"status" validate:"required"`
	Hour        int       `json:"hour" validate:"required"`
	AdditionFee int       `json:"addition_fee" validate:"required"`
	LawyerID    uuid.UUID `json:"lawyer_id" validate:"required"`
	ClientID    uuid.UUID `json:"client_id" validate:"required"`
}

type AddCaseResponseDTO struct {
	Message     string    `json:"message"`
	ID          uuid.UUID `json:"id"`
	Subject     string    `json:"subject"`
	Media       string    `json:"media"`
	Notes       string    `json:"notes"`
	Status      string    `json:"status"`
	Hour        int       `json:"hour"`
	AdditionFee int       `json:"addition_fee"`
	LawyerID    uuid.UUID `json:"lawyer_id"`
	ClientID    uuid.UUID `json:"client_id"`
}

type DeleteCaseRequestDTO struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

type UpdateCaseByIDRequestDTO struct {
	ID          uuid.UUID `json:"id" validate:"required"`
	Subject     string    `json:"subject" validate:"required"`
	Media       string    `json:"media" validate:"required"`
	Notes       string    `json:"notes" validate:"required"`
	Status      string    `json:"status" validate:"required"`
	Hour        int       `json:"hour" validate:"required"`
	AdditionFee int       `json:"addition_fee" validate:"required"`
	LawyerID    uuid.UUID `json:"lawyer_id" validate:"required"`
	ClientID    uuid.UUID `json:"client_id" validate:"required"`
}
