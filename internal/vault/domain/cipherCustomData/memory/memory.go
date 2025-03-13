package memory

import (
	"context"
	"fmt"
	cipher "github.com/Archetarcher/gophkeeper/internal/vault/domain/cipherCustomData"
	"github.com/google/uuid"
	"sync"
)

type Repository struct {
	ciphers map[uuid.UUID]cipher.CipherCustomData
	sync.Mutex
}

func New() *Repository {
	return &Repository{ciphers: make(map[uuid.UUID]cipher.CipherCustomData)}
}

func (r *Repository) Add(ctx context.Context, d *cipher.CipherCustomData) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.ciphers[d.GetId()]; ok {
		return fmt.Errorf("cipher alreafy exists: %w", cipher.ErrFailedToAddCipherCustomData)
	}
	r.ciphers[d.GetId()] = *d
	return nil
}

func (r *Repository) Update(ctx context.Context, d *cipher.CipherCustomData) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.ciphers[d.GetId()]; !ok {
		return fmt.Errorf("cipher does not exists: %w", cipher.ErrFailedToAddCipherCustomData)
	}
	r.ciphers[d.GetId()] = *d
	return nil
}
