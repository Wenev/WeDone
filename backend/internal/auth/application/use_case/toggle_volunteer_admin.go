package usecase

import (
	"context"

	"github.com/Wenev/WeDone/backend/internal/auth/application/dto"
	"github.com/Wenev/WeDone/backend/internal/auth/domain"
)

type ToggleVolunteerAdminUseCase struct {
	repo domain.UserRepository
}

func (uc *ToggleVolunteerAdminUseCase) Execute(ctx context.Context, input dto.ToggleVolunteerAdminInput) (*dto.UserOutput, error) {
	user, err := uc.repo.FindByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	user.IsAdmin = input.IsAdmin
	if err := uc.repo.Update(ctx, user); err != nil {
		return nil, err
	}

	output := dto.ToUserOutput(user)
	return &output, nil
}
