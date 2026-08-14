package dto

import "github.com/google/uuid"


type AssistantSignUpInput struct {
	Email string
	Initial string
	Password string
	ManagerID *uuid.UUID
}

type ManagerSignUpInput struct {
	Email string
	Initial string
	Password string
	GoogleID string
	VolunteerAdmin bool
}

type SignInInput struct {
	Email *string
	Initial *string
	Password string
}