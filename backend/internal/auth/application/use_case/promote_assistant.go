package usecase

import (
	"context"
	"errors"

	"github.com/Wenev/WeDone/backend/internal/auth/application/dto"
	"github.com/Wenev/WeDone/backend/internal/auth/domain"
)

type PromoteAssistantCase struct {
	repo domain.UserRepository
}

func NewPromoteAssistantCase(repo domain.UserRepository) *PromoteAssistantCase {
	return &PromoteAssistantCase{
		repo: repo,
	}
}

func (uc *PromoteAssistantCase) Execute(ctx context.Context, input dto.PromoteAssistantInput) error {
	assistant, err := uc.repo.FindByID(ctx, input.AssistantID)
	if err != nil {
		return err
	}

	if assistant.Role != domain.RoleAssistant {
		return errors.New("user is not an assistant")
	}

	assistant.Role = domain.RoleManager
	assistant.ManagerID = nil

	if err := uc.repo.Update(ctx, assistant); err != nil {
		return err
	}

	return nil
}
