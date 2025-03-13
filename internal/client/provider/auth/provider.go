package auth

import (
	"context"
	"github.com/Archetarcher/gophkeeper/internal/common/provider"
)

type Provider interface {
	SignUp(ctx context.Context, signUp *SignUp) error
	SignIn(ctx context.Context, signIn *SignIn) (*provider.Token, error)
}
