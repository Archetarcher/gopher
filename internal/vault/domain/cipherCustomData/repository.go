package cipher

import (
	"context"
	"github.com/pkg/errors"
)

var (
	ErrCipherCustomDataNotFound       = errors.New("the cipher was not found in the repository")
	ErrFailedToAddCipherCustomData    = errors.New("failed to add the cipher to the repository")
	ErrFailedToUpdateCipherCustomData = errors.New("failed to update the cipher in the repository")
)

type Repository interface {
	Add(ctx context.Context, cipher *CipherCustomData) error
	Update(ctx context.Context, cipher *CipherCustomData) error
}
