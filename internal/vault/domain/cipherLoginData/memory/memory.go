package memory

import (
	"context"
	"fmt"
	cipher "github.com/Archetarcher/gophkeeper/internal/vault/domain/cipherLoginData"
	"github.com/google/uuid"
	"sync"
)

type Repository struct {
	ciphers map[uuid.UUID]cipher.CipherLoginData
	sync.Mutex
}

func New() *Repository {
	return &Repository{ciphers: make(map[uuid.UUID]cipher.CipherLoginData)}
}

func (r *Repository) Add(ctx context.Context, u *cipher.CipherLoginData) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.ciphers[u.GetId()]; ok {
		return fmt.Errorf("cipher alreafy exists: %w", cipher.ErrFailedToAddCipherLoginData)
	}
	r.ciphers[u.GetId()] = *u
	return nil
}

func (r *Repository) Update(ctx context.Context, u *cipher.CipherLoginData) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.ciphers[u.GetId()]; !ok {
		return fmt.Errorf("cipher does not exists: %w", cipher.ErrFailedToAddCipherLoginData)
	}
	r.ciphers[u.GetId()] = *u
	return nil
}
