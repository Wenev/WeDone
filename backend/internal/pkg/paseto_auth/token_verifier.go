package pasetoauth

import (
	"errors"

	"aidanwoods.dev/go-paseto"
	"github.com/google/uuid"
)

var ErrInvalidToken = errors.New("invalid or expired token")

type Verifier struct {
	publicKey paseto.V4AsymmetricPublicKey
}

func NewVerifier(publicKey paseto.V4AsymmetricPublicKey) *Verifier {
	return &Verifier{
		publicKey: publicKey,
	}
}

func (v *Verifier) Verify(tokenString string) (*Claims, error) {
	parser := paseto.NewParser()
	parser.AddRule(paseto.NotExpired())

	token, err := parser.ParseV4Public(v.publicKey, tokenString, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	sub, err := token.GetSubject()
	if err != nil {
		return nil, ErrInvalidToken
	}

	userID, err := uuid.Parse(sub)
	if err != nil {
		return nil, ErrInvalidToken
	}

	role, err := token.GetString("role")
	if err != nil {
		return nil, ErrInvalidToken
	}

	var isAdmin bool
	if err := token.Get("is_admin", &isAdmin); err != nil {
		return nil, ErrInvalidToken
	}

	var managerID *uuid.UUID
	if managerIDStr, err := token.GetString("manager_id"); err == nil && managerIDStr != "" {
		parsed, parseErr := uuid.Parse(managerIDStr)
		if parseErr != nil {
			return nil, ErrInvalidToken
		}
		managerID = &parsed
	}
	return &Claims{UserID: userID, Role: role, IsAdmin: isAdmin, ManagerID: managerID}, nil
}
