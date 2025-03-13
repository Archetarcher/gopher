package secret

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	ErrSecretNotFound       = errors.New("the secret was not found in the repository")
	ErrFailedToAddSecret    = errors.New("failed to add the secret to the repository")
	ErrFailedToUpdateSecret = errors.New("failed to update the secret in the repository")
)

type Repository interface {
	Get(ctx context.Context, uuid uuid.UUID) (*Secret, error)
	GetSecretByUserAndKey(ctx context.Context, userId uuid.UUID, key string) (*Secret, error)
	GetAllSecretByUser(ctx context.Context, userId uuid.UUID) ([]Secret, error)
	Add(ctx context.Context, cipher *Secret) error
	Update(ctx context.Context, cipher *Secret) error
}
