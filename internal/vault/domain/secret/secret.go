package secret

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrInvalidSecret = errors.New("secret has to have  valid fields")
)

// Secret is an aggregate for auth
type Secret struct {
	id         uuid.UUID
	userId     uuid.UUID
	cipherType Type
	data       []byte
	key        []byte
	deletedAt  time.Time
	createdAt  time.Time
	updatedAt  time.Time
}

// NewSecret is a Factory to create a new Secret aggregate
// It will validate that the data, key, userId, cipherType are not empty
func NewSecret(data, key []byte, cipherType string, userId uuid.UUID) (*Secret, error) {
	if len(data) == 0 {
		return nil, errors.Wrap(ErrInvalidSecret, "data does not provided")
	}
	if len(key) == 0 {
		return nil, errors.Wrap(ErrInvalidSecret, "key does not provided")
	}
	if cipherType == "" {
		return nil, errors.Wrap(ErrInvalidSecret, "secret type does not provided")
	}
	if userId == uuid.Nil {
		return nil, errors.Wrap(ErrInvalidSecret, "incorrect userId")
	}
	ct, err := NewSecretTypeFromString(cipherType)
	if err != nil {
		return nil, errors.Wrap(ErrInvalidSecret, err.Error())
	}
	return &Secret{
		id:         uuid.New(),
		userId:     userId,
		cipherType: ct,
		data:       data,
		key:        key,
		createdAt:  time.Now(),
		updatedAt:  time.Now(),
	}, nil
}

func (u *Secret) GetId() uuid.UUID {
	return u.id
}
func (u *Secret) GetUserId() uuid.UUID {
	return u.userId
}
func (u *Secret) GetKey() []byte {
	return u.key
}
func (u *Secret) GetData() []byte {
	return u.data
}
func (u *Secret) GetType() string {
	return u.cipherType.s
}
func (u *Secret) GetCreatedAt() time.Time {
	return u.createdAt
}
func (u *Secret) GetUpdatedAt() time.Time {
	return u.updatedAt
}
