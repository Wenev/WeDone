package utility

import "context"

type GoogleUserInfo struct {

}

type OAuthProvider interface {
	AuthCodeUrl(state string) string
	Exchange(ctx context.Context, code string) (*GoogleUserInfo, error)
}