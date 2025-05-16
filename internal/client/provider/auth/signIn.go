package auth

import (
	"github.com/pkg/errors"
)

var (
	ErrInvalidSignInData = errors.New("invalid sign in data")
)

// SignIn is an aggregate for auth
type SignIn struct {
	login    string
	password string
}

// NewSignIn is a Factory to create a new SignIn aggregate
// It will validate that the login, password, firstname, lastname are not empty
func NewSignIn(login, password string) (*SignIn, error) {
	if login == "" {
		return nil, errors.Wrap(ErrInvalidSignInData, "empty login")
	}
	if password == "" {
		return nil, errors.Wrap(ErrInvalidSignInData, "empty password")
	}
	return &SignIn{
		login:    login,
		password: password,
	}, nil
}

func (u *SignIn) GetLogin() string {
	return u.login
}
func (u *SignIn) GetPassword() string {
	return u.password
}
