package dto

import "github.com/google/uuid"

type AddSpecialtiesRequestDTO struct {
	Name string `json:"name" validate:"required"`
}

type AddSpecialtiesResponseDTO struct {
	Message string    `json:"message"`
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
}

type AddSpecialtiesToLawyerRequestDTO struct {
	LawyerID      uuid.UUID `json:"lawyer_id" validate:"required"`
	SpecialtiesID uuid.UUID `json:"specialties_id" validate:"required"`
}

type AddSpecialtiesToLawyerResponseDTO struct {
	Message       string    `json:"message"`
	LawyerID      uuid.UUID `json:"lawyer_id"`
	SpecialtiesID uuid.UUID `json:"specialties_id"`
}
