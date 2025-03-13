package memory

import (
	"context"
	"fmt"
	cipher "github.com/Archetarcher/gophkeeper/internal/vault/domain/cipherCustomBinaryData"
	"github.com/google/uuid"
	"sync"
)

type Repository struct {
	ciphers map[uuid.UUID]cipher.CipherCustomBinaryData
	sync.Mutex
}

func New() *Repository {
	return &Repository{ciphers: make(map[uuid.UUID]cipher.CipherCustomBinaryData)}
}

func (r *Repository) Add(ctx context.Context, d *cipher.CipherCustomBinaryData) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.ciphers[d.GetId()]; ok {
		return fmt.Errorf("cipher alreafy exists: %w", cipher.ErrFailedToAddCipherCustomBinaryData)
	}
	r.ciphers[d.GetId()] = *d
	return nil
}

func (r *Repository) Update(ctx context.Context, d *cipher.CipherCustomBinaryData) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.ciphers[d.GetId()]; !ok {
		return fmt.Errorf("cipher does not exists: %w", cipher.ErrFailedToAddCipherCustomBinaryData)
	}
	r.ciphers[d.GetId()] = *d
	return nil
}
