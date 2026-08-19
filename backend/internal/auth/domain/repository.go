package domain

import (
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByGoogleID(ctx context.Context, googleID string) (*User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*User, error)
	FindByInitial(ctx context.Context, initial string) (*User, error)
	Update(ctx context.Context, user *User) error
	PromoteToManager(ctx context.Context, id uuid.UUID) error
	UpdateAdminStatus(ctx context.Context, id uuid.UUID, isAdmin bool) error
	UpdatePassword(ctx context.Context, id uuid.UUID, password string) error
}

type CredentialLookupRepository interface {
	FindByEmailForAuthentication(ctx context.Context, email string) (*User, error)
	FindByInitialForAuthentication(ctx context.Context, initial string) (*User, error)
}
