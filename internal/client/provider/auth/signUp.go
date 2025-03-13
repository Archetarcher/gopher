package auth

import (
	"github.com/pkg/errors"
)

var (
	ErrInvalidSignUpData = errors.New("invalid sign up data")
)

// SignUp is an aggregate for auth
type SignUp struct {
	login     string
	password  string
	firstname string
	lastname  string
}

// NewSignUp is a Factory to create a new SignIn aggregate
// It will validate that the login, password, firstname, lastname are not empty
func NewSignUp(login, password, firstname, lastname string) (*SignUp, error) {
	if login == "" {
		return nil, errors.Wrap(ErrInvalidSignInData, "empty login")
	}
	if password == "" {
		return nil, errors.Wrap(ErrInvalidSignInData, "empty password")
	}
	if firstname == "" {
		return nil, errors.Wrap(ErrInvalidSignInData, "empty firstname")
	}
	if lastname == "" {
		return nil, errors.Wrap(ErrInvalidSignInData, "empty lastname")
	}
	return &SignUp{
		login:     login,
		password:  password,
		lastname:  lastname,
		firstname: firstname,
	}, nil
}
