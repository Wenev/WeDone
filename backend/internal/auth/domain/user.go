package domain

import (
	"errors"

	"github.com/google/uuid"
)

type Role string

const (
	RoleManager Role = "manager"
	RoleAssistant Role = "assistant"
)

var (
	ErrManagerRequiresNoParent = errors.New("A manager must not have a manager_id")
	ErrInvalidEmail = errors.New("email is invalid or empty")
)

type User struct {
	ID uuid.UUID
	Email string
	Initial string
	Role Role
	ManagerID *uuid.UUID
	IsAdmin bool
	Password *string
	GoogleID *string
}

func NewAssistant(email, initial string, managerId *uuid.UUID) (*User, error) {
	if email == "" {
		return nil, ErrInvalidEmail
	}
	return &User{
		ID: uuid.New(),
		Email: email,
		Initial: initial,
		ManagerID: managerId,
		Role: RoleAssistant,
		IsAdmin: false,
	}, nil
}

func NewManager(email, initial, googleId string) (*User, error) {
	if email == "" {
		return nil, ErrInvalidEmail
	}

	var googleIDptr *string
	if googleId != "" {
		googleIDptr = &googleId
	}

	return &User{
		ID: uuid.New(),
		Email: email,
		Initial: initial,
		GoogleID: googleIDptr,
		ManagerID: nil,
		Role: RoleManager,
		IsAdmin: false,
	}, nil
}