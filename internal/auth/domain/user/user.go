package user

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrInvalidPerson = errors.New("auth has to have a valid person")
)

// User is an aggregate for auth
type User struct {
	id        uuid.UUID
	firstname string
	lastname  string
	login     string
	hash      string
	createdAt time.Time
	updatedAt time.Time
}

// NewUser is a Factory to create a new User aggregate
// It will validate that the login, password, firstname, lastname are not empty
func NewUser(login, hash, firstname, lastname string) (*User, error) {
	if login == "" {
		return nil, errors.Wrap(ErrInvalidPerson, "empty login")
	}
	if hash == "" {
		return nil, errors.Wrap(ErrInvalidPerson, "empty password")
	}
	if firstname == "" {
		return nil, errors.Wrap(ErrInvalidPerson, "empty firstname")
	}
	if lastname == "" {
		return nil, errors.Wrap(ErrInvalidPerson, "empty lastname")
	}
	return &User{
		id:        uuid.New(),
		firstname: firstname,
		lastname:  lastname,
		login:     login,
		hash:      hash,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

// UnmarshalUserFromDatabase marshals db user to domain user
// It's not constructor, use only for db unmarshalling
func UnmarshalUserFromDatabase(id uuid.UUID, login, hash, firstname, lastname string, createdAt time.Time, updatedAt time.Time) (*User, error) {
	u, err := NewUser(login, hash, firstname, lastname)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal db user model")
	}
	u.id = (id)
	u.createdAt = createdAt
	u.updatedAt = updatedAt
	return u, nil
}
func (u *User) GetLogin() string {
	return u.login
}
func (u *User) GetId() uuid.UUID {
	return u.id
}
func (u *User) GetFirstname() string {
	return u.firstname
}
func (u *User) GetLastname() string {
	return u.lastname
}
func (u *User) GetHash() string {
	return u.hash
}
