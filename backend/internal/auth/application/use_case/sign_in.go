package usecase

import (
	"context"
	"errors"
	"strings"

	"github.com/Wenev/WeDone/backend/internal/auth/application/dto"
	"github.com/Wenev/WeDone/backend/internal/auth/application/utility"
	"github.com/Wenev/WeDone/backend/internal/auth/domain"
	pasetoauth "github.com/Wenev/WeDone/backend/internal/pkg/paseto_auth"
)

type SignInUseCase struct {
	repo   domain.UserRepository
	hasher utility.PasswordHasher
	issuer pasetoauth.Issuer
}

func NewSignInUseCase(repo domain.UserRepository, hasher utility.PasswordHasher, issuer pasetoauth.Issuer) *SignInUseCase {
	return &SignInUseCase{
		repo:   repo,
		hasher: hasher,
		issuer: issuer,
	}
}

func (uc *SignInUseCase) Execute(ctx context.Context, input dto.SignInInput) (*dto.AuthOutput, error) {
	identifier := strings.TrimSpace(input.Identifier)
	if identifier == "" {
		return nil, errors.New("email or initial is required")
	}

	var foundUser *domain.User
	var err error
	if strings.Contains(identifier, "@") {
		foundUser, err = uc.repo.FindByEmail(ctx, identifier)
	} else {
		foundUser, err = uc.repo.FindByInitial(ctx, identifier)
	}
	if err != nil && !errors.Is(err, domain.ErrUserNotFound) {
		return nil, err
	}
	if foundUser == nil {
		return nil, domain.ErrInvalidCredentials
	}

	
	if passValidation, err := uc.hasher.Compare(foundUser.Password, input.Password); err != nil || !passValidation {
		return nil, domain.ErrInvalidCredentials
	}

	token, err := uc.issuer.Issue(foundUser.ID, string(foundUser.Role), foundUser.IsAdmin, foundUser.ManagerID)
	if err != nil {
		return nil, err
	}

	output := dto.ToUserOutput(foundUser)

	return &dto.AuthOutput{
		Token: token,
		User:  output,
	}, nil
}
