package cipher

import (
	"github.com/Archetarcher/gophkeeper/internal/vault"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrInvalidCipherCustomData = errors.New("cipher has to have  valid fields")
)

// CipherCustomData is an aggregate
type CipherCustomData struct {
	cipher *vault.Cipher

	key   []byte
	value []byte
}

// NewCipherCustomData is a Factory to create a new CipherCustomData aggregate
// It will validate that the data, key, userId, cipherType are not empty
func NewCipherCustomData(key, value, meta []byte, userId uuid.UUID) (*CipherCustomData, error) {
	if len(key) == 0 {
		return nil, errors.Wrap(ErrInvalidCipherCustomData, "key does not provided")
	}
	if len(value) == 0 {
		return nil, errors.Wrap(ErrInvalidCipherCustomData, "value does not provided")
	}
	if userId == uuid.Nil {
		return nil, errors.Wrap(ErrInvalidCipherCustomData, "incorrect userId")
	}

	return &CipherCustomData{
		cipher: &vault.Cipher{
			Id:        uuid.New(),
			UserId:    userId,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			MetaData:  meta,
		},
		key:   key,
		value: value,
	}, nil
}

// UnmarshalCipherCustomDataFromDatabase marshals db model to domain aggregate
// It's not constructor, use only for db unmarshalling
func UnmarshalCipherCustomDataFromDatabase(id uuid.UUID, key, value, meta []byte, userId uuid.UUID, createdAt, updatedAt, deletedAt time.Time) (*CipherCustomData, error) {
	c, err := NewCipherCustomData(key, value, meta, userId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal db model")
	}
	c.cipher.Id = id
	c.cipher.CreatedAt = createdAt
	c.cipher.UpdatedAt = updatedAt
	c.cipher.DeletedAt = deletedAt
	return c, nil
}
func (u *CipherCustomData) GetId() uuid.UUID {
	return u.cipher.Id
}
func (u *CipherCustomData) GetKey() []byte {
	return u.key
}
func (u *CipherCustomData) GetValue() []byte {
	return u.value
}
func (u *CipherCustomData) GetMetaData() []byte {
	return u.cipher.MetaData
}
func (u *CipherCustomData) GetUserId() uuid.UUID {
	return u.cipher.UserId
}
func (u *CipherCustomData) GetCreatedAt() time.Time {
	return u.cipher.CreatedAt
}
func (u *CipherCustomData) GetUpdatedAt() time.Time {
	return u.cipher.UpdatedAt
}
