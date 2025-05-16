package cipher

import (
	"github.com/Archetarcher/gophkeeper/internal/vault"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrInvalidCipherLoginData = errors.New("secret has to have  valid fields")
)

// CipherLoginData is an aggregate for auth
type CipherLoginData struct {
	cipher *vault.Cipher

	uri      []byte
	login    []byte
	password []byte
}

// NewCipherLoginData is a Factory to create a new CipherLoginData aggregate
// It will validate that the data, key, userId, cipherType are not empty
func NewCipherLoginData(uri, login, password, meta []byte, userId uuid.UUID) (*CipherLoginData, error) {
	if len(uri) == 0 {
		return nil, errors.Wrap(ErrInvalidCipherLoginData, "uri does not provided")
	}
	if len(login) == 0 {
		return nil, errors.Wrap(ErrInvalidCipherLoginData, "login does not provided")
	}
	if len(password) == 0 {
		return nil, errors.Wrap(ErrInvalidCipherLoginData, "password does not provided")
	}
	if userId == uuid.Nil {
		return nil, errors.Wrap(ErrInvalidCipherLoginData, "incorrect userId")
	}

	return &CipherLoginData{
		cipher: &vault.Cipher{
			Id:        uuid.New(),
			UserId:    userId,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			MetaData:  meta,
		},
		uri:      uri,
		login:    login,
		password: password,
	}, nil
}

// UnmarshalCipherLoginDataFromDatabase marshals db model to domain aggregate
// It's not constructor, use only for db unmarshalling
func UnmarshalCipherLoginDataFromDatabase(id uuid.UUID, uri, login, password, meta []byte, userId uuid.UUID, createdAt, updatedAt, deletedAt time.Time) (*CipherLoginData, error) {
	c, err := NewCipherLoginData(uri, login, password, meta, userId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal db model")
	}
	c.cipher.Id = id
	c.cipher.CreatedAt = createdAt
	c.cipher.UpdatedAt = updatedAt
	c.cipher.DeletedAt = deletedAt
	return c, nil
}
func (u *CipherLoginData) GetId() uuid.UUID {
	return u.cipher.Id
}
func (u *CipherLoginData) GetUserId() uuid.UUID {
	return u.cipher.UserId
}
func (u *CipherLoginData) GetUri() []byte {
	return u.uri
}
func (u *CipherLoginData) GetLogin() []byte {
	return u.login
}
func (u *CipherLoginData) GetPassword() []byte {
	return u.password
}
func (u *CipherLoginData) GetMetaData() []byte {
	return u.cipher.MetaData
}
func (u *CipherLoginData) GetCreatedAt() time.Time {
	return u.cipher.CreatedAt
}
func (u *CipherLoginData) GetUpdatedAt() time.Time {
	return u.cipher.UpdatedAt
}
