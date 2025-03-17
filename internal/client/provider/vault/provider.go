package vault

import (
	"context"
	"github.com/pkg/errors"
)

var (
	ErrTokenExpired           = errors.New("jwt token is expired, please sign in first")
	ErrFailedToRememberCipher = errors.New("failed to remember cipher ")
)

type Provider interface {
	RememberCipherLogin(ctx context.Context, cipher *RememberCipherLoginData) error
	RememberCipherCustom(ctx context.Context, cipher *RememberCipherCustomData) error
	RememberCipherCustomBinary(ctx context.Context, cipher *RememberCipherCustomBinaryData) error
	RememberCipherCard(ctx context.Context, cipher *RememberCipherCardData) error
}
