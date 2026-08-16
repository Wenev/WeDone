package utility

import "context"

type GoogleUserInfo struct {
	GoogleID      string
	Email         string
	VerifiedEmail bool
}

type OAuthProvider interface {
	AuthCodeUrl(state string) string
	Exchange(ctx context.Context, code string) (*GoogleUserInfo, error)
}
