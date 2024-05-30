package dto

import "github.com/google/uuid"

type LoginRequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponseDTO struct {
	Message  string    `json:"message"`
	Token    string    `json:"token"`
	Username string    `json:"username"`
	UserID   uuid.UUID `json:"user_id"`
}
