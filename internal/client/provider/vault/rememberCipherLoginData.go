package vault

import (
	"github.com/pkg/errors"
)

var (
	ErrInvalidCipherLoginData = errors.New("cipher has to have  valid fields")
)

// RememberCipherLoginData is an aggregate for auth
type RememberCipherLoginData struct {
	uri      string
	login    string
	password string
	meta     string
}

// NewRememberCipherLoginData is a Factory to create a new CipherLoginData aggregate
// It will validate that the data, key, userId, cipherType are not empty
func NewRememberCipherLoginData(uri, login, password, meta string) (*RememberCipherLoginData, error) {
	if uri == "" {
		return nil, errors.Wrap(ErrInvalidCipherLoginData, "uri does not provided")
	}
	if login == "" {
		return nil, errors.Wrap(ErrInvalidCipherLoginData, "login does not provided")
	}
	if password == "" {
		return nil, errors.Wrap(ErrInvalidCipherLoginData, "password does not provided")
	}

	return &RememberCipherLoginData{
		uri:      uri,
		login:    login,
		password: password,
		meta:     meta,
	}, nil
}

func (u *RememberCipherLoginData) GetUri() string {
	return u.uri
}
func (u *RememberCipherLoginData) GetLogin() string {
	return u.login
}
func (u *RememberCipherLoginData) GetPassword() string {
	return u.password
}
func (u *RememberCipherLoginData) GetMeta() string {
	return u.meta
}
