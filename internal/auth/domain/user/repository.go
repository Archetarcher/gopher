package user

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrUserAlreadyRegistered = errors.New("the user with this login already exists")
	ErrUserNotFound          = errors.New("the user was not found in the repository")
	ErrFailedToAddUser       = errors.New("failed to add the user to the repository")
	ErrFailedToUpdateUser    = errors.New("failed to update the user in the repository")
	ErrSignIn                = errors.New("failed to sign in")
)

type Repository interface {
	GetByLogin(ctx context.Context, login string) (*User, error)
	Get(ctx context.Context, uuid uuid.UUID) (*User, error)
	Add(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
}
