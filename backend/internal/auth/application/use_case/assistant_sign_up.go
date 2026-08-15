package usecase

import (
	"context"
	"errors"

	"github.com/Wenev/WeDone/backend/internal/auth/application/dto"
	"github.com/Wenev/WeDone/backend/internal/auth/application/utility"
	"github.com/Wenev/WeDone/backend/internal/auth/domain"
)

var (
	ErrEmailTaken    = errors.New("an account with this email already exists")
	ErrInitialTaken  = errors.New("an account with this Initial already exists")
	ErrGoogleIDTaken = errors.New("an account with this Google Acount already exists")
)

type AssistantSignUpUseCase struct {
	repo   domain.UserRepository
	hasher utility.PasswordHasher
}

func NewAssistantSignUpUseCase(repo domain.UserRepository, hasher utility.PasswordHasher) *AssistantSignUpUseCase {
	return &AssistantSignUpUseCase{
		repo:   repo,
		hasher: hasher,
	}
}

func (uc *AssistantSignUpUseCase) Execute(ctx context.Context, input dto.AssistantSignUpInput) (*dto.UserOutput, error) {
	_, err := uc.repo.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, ErrEmailTaken
	}
	_, err = uc.repo.FindByInitial(ctx, input.Initial)
	if err != nil {
		return nil, ErrInitialTaken
	}

	hashed, err := uc.hasher.Hash(input.Password)
	if err != nil {
		return nil, err
	}

	user, err := domain.NewAssistant(input.Email, input.Initial, input.ManagerID)
	if err != nil {
		return nil, err
	}
	user.Password = &hashed

	if err := uc.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	output := dto.ToUserOutput(user)
	return &output, nil
}
