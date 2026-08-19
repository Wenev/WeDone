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
	ErrOnlyManagerCanBeAdmin   = errors.New("only a manager can be an admin")
	ErrInvalidInitial          = errors.New("initial is invalid or empty")
)

type PublicUser struct {
	ID        uuid.UUID
	Email     string
	Initial   string
	GoogleID  string
	Role      Role
	ManagerID *uuid.UUID
	IsAdmin   bool
}

type User struct {
	PublicUser
	Password string
}

func (u *User) Validate() error {
	if _, err := NormalizeEmail(u.Email); err != nil {
		return ErrInvalidEmail
	}
	if _, err := NormalizeInitial(u.Initial); err != nil {
		return ErrInvalidInitial
	}
	if u.Role == RoleManager && u.ManagerID != nil {
		return ErrManagerRequiresNoParent
	}
	return nil
}

func NewAssistant(email, initial string, managerId *uuid.UUID) (*User, error) {
	normalizedEmail, err := NormalizeEmail(email)
	if err != nil {
		return nil, ErrInvalidEmail
	}
	normalizedInitial, err := NormalizeInitial(initial)
	if err != nil {
		return nil, ErrInvalidInitial
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &User{
		PublicUser: PublicUser{
			ID:        id,
			Email:     normalizedEmail,
			Initial:   normalizedInitial,
			ManagerID: managerId,
			Role:      RoleAssistant,
			IsAdmin:   false,
		},
	}, nil
}

func NewManager(email, initial, googleId string) (*User, error) {
	normalizedEmail, err := NormalizeEmail(email)
	if err != nil {
		return nil, ErrInvalidEmail
	}
	normalizedInitial, err := NormalizeInitial(initial)
	if err != nil {
		return nil, ErrInvalidInitial
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &User{
		PublicUser: PublicUser{
			ID:        id,
			Email:     normalizedEmail,
			Initial:   normalizedInitial,
			GoogleID:  googleId,
			ManagerID: nil,
			Role:      RoleManager,
			IsAdmin:   false,
		},
	}, nil
}

func (u *User) SetIsAdmin(isAdmin bool) error {
	if isAdmin && u.Role != RoleManager {
		return ErrOnlyManagerCanBeAdmin
	}
	u.IsAdmin = isAdmin
	return nil
}
