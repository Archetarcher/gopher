package memory

import (
	"context"
	"fmt"
	"github.com/Archetarcher/gophkeeper/internal/vault/domain/secret"
	"github.com/google/uuid"
	"sync"
)

type Repository struct {
	secrets map[uuid.UUID]secret.Secret
	sync.Mutex
}

func New() *Repository {
	return &Repository{secrets: make(map[uuid.UUID]secret.Secret)}
}

func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*secret.Secret, error) {
	if u, ok := r.secrets[id]; ok {
		return &u, nil
	}
	return nil, secret.ErrSecretNotFound
}
func (r *Repository) GetSecretByUserAndKey(ctx context.Context, userId uuid.UUID, key string) (*secret.Secret, error) {
	for _, ru := range r.secrets {
		if ru.GetUserId() == userId && string(ru.GetKey()) == key {
			return &ru, nil
		}
	}
	return nil, secret.ErrSecretNotFound
}
func (r *Repository) GetAllSecretByUser(ctx context.Context, userId uuid.UUID) ([]secret.Secret, error) {
	var secrets []secret.Secret
	for _, ru := range r.secrets {
		if ru.GetUserId() == userId {
			secrets = append(secrets, ru)
		}
	}
	if len(secrets) == 0 {
		return nil, secret.ErrSecretNotFound

	}
	return secrets, nil
}

func (r *Repository) Add(ctx context.Context, u *secret.Secret) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.secrets[u.GetId()]; ok {
		return fmt.Errorf("secret alreafy exists: %w", secret.ErrFailedToAddSecret)
	}
	r.secrets[u.GetId()] = *u
	return nil
}

func (r *Repository) Update(ctx context.Context, u *secret.Secret) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.secrets[u.GetId()]; !ok {
		return fmt.Errorf("secret does not exists: %w", secret.ErrFailedToAddSecret)
	}
	r.secrets[u.GetId()] = *u
	return nil
}
