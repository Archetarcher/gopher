package memory

import (
	"context"
	"fmt"
	cipher "github.com/Archetarcher/gophkeeper/internal/vault/domain/cipherCardData"
	"github.com/google/uuid"
	"sync"
)

type Repository struct {
	ciphers map[uuid.UUID]cipher.CipherCardData
	sync.Mutex
}

func New() *Repository {
	return &Repository{ciphers: make(map[uuid.UUID]cipher.CipherCardData)}
}

func (r *Repository) Add(ctx context.Context, d *cipher.CipherCardData) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.ciphers[d.GetId()]; ok {
		return fmt.Errorf("cipher alreafy exists: %w", cipher.ErrFailedToAddCipherCardData)
	}
	r.ciphers[d.GetId()] = *d
	return nil
}

func (r *Repository) Update(ctx context.Context, d *cipher.CipherCardData) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.ciphers[d.GetId()]; !ok {
		return fmt.Errorf("cipher does not exists: %w", cipher.ErrFailedToAddCipherCardData)
	}
	r.ciphers[d.GetId()] = *d
	return nil
}
