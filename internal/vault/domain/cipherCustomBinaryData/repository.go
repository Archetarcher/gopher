package cipher

import (
	"context"
	"github.com/pkg/errors"
)

var (
	ErrCipherCustomBinaryDataNotFound       = errors.New("the cipher was not found in the repository")
	ErrFailedToAddCipherCustomBinaryData    = errors.New("failed to add the cipher to the repository")
	ErrFailedToUpdateCipherCustomBinaryData = errors.New("failed to update the cipher in the repository")
)

type Repository interface {
	Add(ctx context.Context, cipher *CipherCustomBinaryData) error
	Update(ctx context.Context, cipher *CipherCustomBinaryData) error
}
