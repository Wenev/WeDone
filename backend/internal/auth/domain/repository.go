package domain

import (
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByProviderID(ctx context.Context, providerId string, provider string) (*User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*User, error)
	FindByInitial(ctx context.Context, initial string) (*User, error)
	Update(ctx context.Context, user *User) error
}