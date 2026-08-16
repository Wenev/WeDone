package usecase

import (
	"context"

	"github.com/Wenev/WeDone/backend/internal/auth/application/dto"
	"github.com/Wenev/WeDone/backend/internal/auth/application/utility"
	"github.com/Wenev/WeDone/backend/internal/auth/domain"
)



type ManagerSignUpUseCase struct {
	repo domain.UserRepository
	hasher utility.PasswordHasher
}

func (uc *ManagerSignUpUseCase) Execute(ctx context.Context, input dto.ManagerSignUpInput) (*dto.UserOutput, error) {
	_, err := uc.repo.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, ErrEmailTaken
	}
	_, err = uc.repo.FindByInitial(ctx, input.Initial)
	if err != nil {
		return nil, ErrInitialTaken
	}
	_, err = uc.repo.FindByProviderID(ctx, input.GoogleID, "google")
	if err != nil {
		return nil, ErrInitialTaken
	}
	hashed, err := uc.hasher.Hash(input.Password)
	if err != nil {
		return nil, err
	}

	user, err := domain.NewManager(input.Email, input.Initial, input.GoogleID)
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