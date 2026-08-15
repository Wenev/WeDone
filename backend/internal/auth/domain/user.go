package domain

import (
	"errors"

	"github.com/google/uuid"
)

type Role string

const (
	RoleManager   Role = "manager"
	RoleAssistant Role = "assistant"
)

var (
	ErrManagerRequiresNoParent = errors.New("A manager must not have a manager_id")
	ErrInvalidEmail            = errors.New("email is invalid or empty")
	ErrUserNotFound            = errors.New("user not found")
)

type User struct {
	ID        uuid.UUID
	Email     string
	Initial   string
	Password  string
	GoogleID  string
	Role      Role
	ManagerID *uuid.UUID
	IsAdmin   bool
}

func NewAssistant(email, initial string, managerId *uuid.UUID) (*User, error) {
	if email == "" {
		return nil, ErrInvalidEmail
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &User{
		ID:        id,
		Email:     email,
		Initial:   initial,
		ManagerID: managerId,
		Role:      RoleAssistant,
		IsAdmin:   false,
	}, nil
}

func NewManager(email, initial, googleId string) (*User, error) {
	if email == "" {
		return nil, ErrInvalidEmail
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &User{
		ID:        id,
		Email:     email,
		Initial:   initial,
		GoogleID:  googleId,
		ManagerID: nil,
		Role:      RoleManager,
		IsAdmin:   false,
	}, nil
}
