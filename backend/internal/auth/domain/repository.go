package domain

import "context"

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByGoogleID(ctx context.Context, googleId string) (*User, error)
	FindByID(ctx context.Context, googleId string) (*User, error)
	FindByInitial(ctx context.Context, initial string) (*User, error)
}