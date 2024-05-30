package dto

import "github.com/google/uuid"

type UserRegisterRequestDTO struct {
	Username    string `json:"username" validate:"required"`
	FullName    string `json:"full_name" validate:"required"`
	Email       string `json:"email" validate:"required"`
	Password    string `json:"password" validate:"required"`
	BirthDate   string `json:"birth_date" validate:"required"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Bio         string `json:"bio"`
}

type UserRegisterResponseDTO struct {
	Message     string    `json:"message"`
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	BirthDate   string    `json:"birth_date"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	Bio         string    `json:"bio"`
}

type UpdateUserDataRequestDTO struct {
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	BirthDate   string `json:"birth_date"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Bio         string `json:"bio"`
}

type UpdateUserDataResponseDTO struct {
	Message     string `json:"message"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	BirthDate   string `json:"birth_date"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Bio         string `json:"bio"`
}
