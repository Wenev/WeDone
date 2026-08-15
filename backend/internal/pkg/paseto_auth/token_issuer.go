package pasetoauth

import (
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/google/uuid"
)

type Issuer struct {
	secretKey paseto.V4AsymmetricSecretKey
	ttl time.Duration
}

func NewIssuer(secretKey paseto.V4AsymmetricSecretKey, ttl time.Duration) *Issuer {
	return &Issuer{
		secretKey: secretKey,
		ttl: ttl,
	}
}

func (i *Issuer) Issue(userID uuid.UUID, role string, isAdmin bool, managerID *uuid.UUID) (string, error) {
	token := paseto.NewToken()
	token.SetSubject(userID.String())
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(i.ttl))
	token.SetString("role", role)

	if err := token.Set("is_admin", isAdmin); err != nil {
		return "", err
	}
	if managerID != nil {
		token.SetString("manager_id", managerID.String())
	}

	return token.V4Sign
}