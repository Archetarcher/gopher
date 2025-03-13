package vault

import (
	"context"
	"github.com/Archetarcher/gophkeeper/internal/common/encryption"
	"github.com/pkg/errors"
)

var (
	ErrTokenExpired = errors.New("jwt token is expired")
)

type Provider interface {
	RememberCipherLogin(ctx context.Context, cipher *RememberCipherLoginData) error
	RememberCipherCustom(ctx context.Context, cipher *RememberCipherCustomData) error
	RememberCipherCustomBinary(ctx context.Context, cipher *RememberCipherCustomBinaryData) error
	RememberCipherCard(ctx context.Context, cipher *RememberCipherCardData) error
	StartSession(ctx context.Context, enc *encryption.Asymmetric) error
}
