package cipher

import (
	"context"
	"github.com/pkg/errors"
)

var (
	ErrCipherCardDataNotFound       = errors.New("the cipher was not found in the repository")
	ErrFailedToAddCipherCardData    = errors.New("failed to add the cipher to the repository")
	ErrFailedToUpdateCipherCardData = errors.New("failed to update the cipher in the repository")
)

type Repository interface {
	Add(ctx context.Context, cipher *CipherCardData) error
	Update(ctx context.Context, cipher *CipherCardData) error
}
