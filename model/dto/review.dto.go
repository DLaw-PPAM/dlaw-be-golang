package dto

import "github.com/google/uuid"

type AddReviewRequestDTO struct {
	Rating      int       `json:"rating" validate:"required"`
	Description string    `json:"description" validate:"required"`
	LawyerID    uuid.UUID `json:"lawyer_id" validate:"required"`
	ClientID    uuid.UUID `json:"client_id" validate:"required"`
	ClientName  string    `json:"client_name" validate:"required"`
}

type AddReviewResponseDTO struct {
	Message     string    `json:"message"`
	ID          uuid.UUID `json:"id"`
	Rating      int       `json:"rating"`
	Description string    `json:"description"`
	LawyerID    uuid.UUID `json:"lawyer_id"`
	ClientID    uuid.UUID `json:"client_id"`
	ClientName  string    `json:"client_name"`
}
