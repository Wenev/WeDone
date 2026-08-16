package dto

import (
	"github.com/google/uuid"

	"github.com/Wenev/WeDone/backend/internal/auth/domain"
)

type UserOutput struct {
	ID        uuid.UUID
	Email     string
	Initial   string
	Role      string
	IsAdmin   bool
	ManagerID *uuid.UUID
}

func ToUserOutput(u *domain.User) UserOutput {
	return UserOutput{
		ID:        u.ID,
		Email:     u.Email,
		Initial:   u.Initial,
		Role:      string(u.Role),
		IsAdmin:   u.IsAdmin,
		ManagerID: u.ManagerID,
	}
}

type AuthOutput struct {
	Token string
	User  UserOutput
}
