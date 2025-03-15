package auth

import (
	"context"
	"errors"
	"github.com/Archetarcher/gophkeeper/internal/common/provider"
)

var (
	ErrSignIn = errors.New("failed to sign in user")
	ErrSignUp = errors.New("failed to sign up user")
)

type Provider interface {
	SignUp(ctx context.Context, signUp *SignUp) error
	SignIn(ctx context.Context, signIn *SignIn) (*provider.Token, error)
}
