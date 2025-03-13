package cipher

import (
	"github.com/Archetarcher/gophkeeper/internal/vault"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrInvalidCipherCustomBinaryData = errors.New("cipher has to have  valid fields")
)

// CipherCustomBinaryData is an aggregate
type CipherCustomBinaryData struct {
	cipher *vault.Cipher

	key   []byte
	value []byte
}

// NewCipherCustomBinaryData is a Factory to create a new CipherCustomBinaryData aggregate
// It will validate that the data, key, userId, cipherType are not empty
func NewCipherCustomBinaryData(key, value, meta []byte, userId uuid.UUID) (*CipherCustomBinaryData, error) {
	if len(key) == 0 {
		return nil, errors.Wrap(ErrInvalidCipherCustomBinaryData, "key does not provided")
	}
	if len(value) == 0 {
		return nil, errors.Wrap(ErrInvalidCipherCustomBinaryData, "value does not provided")
	}
	if userId == uuid.Nil {
		return nil, errors.Wrap(ErrInvalidCipherCustomBinaryData, "incorrect userId")
	}

	return &CipherCustomBinaryData{
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

// UnmarshalCipherCustomBinaryDataFromDatabase marshals db model to domain aggregate
// It's not constructor, use only for db unmarshalling
func UnmarshalCipherCustomBinaryDataFromDatabase(id uuid.UUID, key, value, meta []byte, userId uuid.UUID, createdAt, updatedAt, deletedAt time.Time) (*CipherCustomBinaryData, error) {
	c, err := NewCipherCustomBinaryData(key, value, meta, userId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal db model")
	}
	c.cipher.Id = id
	c.cipher.CreatedAt = createdAt
	c.cipher.UpdatedAt = updatedAt
	c.cipher.DeletedAt = deletedAt
	return c, nil
}

func (u *CipherCustomBinaryData) GetId() uuid.UUID {
	return u.cipher.Id
}
func (u *CipherCustomBinaryData) GetKey() []byte {
	return u.key
}
func (u *CipherCustomBinaryData) GetValue() []byte {
	return u.value
}
func (u *CipherCustomBinaryData) GetUserId() uuid.UUID {
	return u.cipher.UserId
}
func (u *CipherCustomBinaryData) GetCreatedAt() time.Time {
	return u.cipher.CreatedAt
}
func (u *CipherCustomBinaryData) GetUpdatedAt() time.Time {
	return u.cipher.UpdatedAt
}
func (u *CipherCustomBinaryData) GetMetaData() []byte {
	return u.cipher.MetaData
}
