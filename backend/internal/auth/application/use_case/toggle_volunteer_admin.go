package usecase

import (
	"context"

	"github.com/Wenev/WeDone/backend/internal/auth/application/dto"
	"github.com/Wenev/WeDone/backend/internal/auth/domain"
)

type ToggleVolunteerAdminUseCase struct {
	repo domain.UserRepository
}

func NewToggleVolunteerAdminUseCase(repo domain.UserRepository) *ToggleVolunteerAdminUseCase {
	return &ToggleVolunteerAdminUseCase{
		repo: repo,
	}
}

func (uc *ToggleVolunteerAdminUseCase) Execute(ctx context.Context, input dto.ToggleVolunteerAdminInput) (*dto.UserOutput, error) {
	user, err := uc.repo.FindByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	if err := user.SetIsAdmin(input.IsAdmin); err != nil {
		return nil, err
	}

	if err := uc.repo.UpdateAdminStatus(ctx, user.ID, user.IsAdmin); err != nil {
		return nil, err
	}

	output := dto.ToUserOutput(user)
	return &output, nil
}
