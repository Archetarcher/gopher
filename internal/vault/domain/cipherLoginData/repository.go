package cipher

import (
	"context"
	"github.com/pkg/errors"
)

var (
	ErrCipherLoginDataNotFound       = errors.New("the cipher was not found in the repository")
	ErrFailedToAddCipherLoginData    = errors.New("failed to add the cipher to the repository")
	ErrFailedToUpdateCipherLoginData = errors.New("failed to update the cipher in the repository")
)

type Repository interface {
	Add(ctx context.Context, cipher *CipherLoginData) error
	Update(ctx context.Context, cipher *CipherLoginData) error
}
