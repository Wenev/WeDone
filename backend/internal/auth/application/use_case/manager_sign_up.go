package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/Wenev/WeDone/backend/internal/auth/application/dto"
	"github.com/Wenev/WeDone/backend/internal/auth/application/utility"
	"github.com/Wenev/WeDone/backend/internal/auth/domain"
)

type ManagerSignUpUseCase struct {
	repo   domain.UserRepository
	hasher utility.PasswordHasher
}

func NewManagerSignUpUseCase(repo domain.UserRepository, hasher utility.PasswordHasher) *ManagerSignUpUseCase {
	return &ManagerSignUpUseCase{
		repo:   repo,
		hasher: hasher,
	}
}

func (uc *ManagerSignUpUseCase) Execute(ctx context.Context, input dto.ManagerSignUpInput) (*dto.UserOutput, error) {
	existing, err := uc.repo.FindByEmail(ctx, input.Email)
	if err == nil && existing != nil {
		return nil, ErrEmailTaken
	}
	if err != nil && !errors.Is(err, domain.ErrUserNotFound) {
		return nil, fmt.Errorf("failed to verify email: %w", err)
	}

	existing, err = uc.repo.FindByInitial(ctx, input.Initial)
	if err == nil && existing != nil {
		return nil, ErrInitialTaken
	}
	if err != nil && !errors.Is(err, domain.ErrUserNotFound) {
		return nil, fmt.Errorf("failed to verify initial: %w", err)
	}

	if input.GoogleID != "" {
		existing, err = uc.repo.FindByGoogleID(ctx, input.GoogleID)
		if err == nil && existing != nil {
			return nil, ErrGoogleIDTaken
		}
		if err != nil && !errors.Is(err, domain.ErrUserNotFound) {
			return nil, fmt.Errorf("failed to verify provider ID: %w", err)
		}
	}

	hashed, err := uc.hasher.Hash(input.Password)
	if err != nil {
		return nil, err
	}

	user, err := domain.NewManager(input.Email, input.Initial, input.GoogleID)
	if err != nil {
		return nil, err
	}
	user.Password = hashed

	if err := uc.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	output := dto.ToUserOutput(user)
	return &output, nil
}