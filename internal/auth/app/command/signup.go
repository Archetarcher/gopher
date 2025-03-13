package command

import (
	"context"
	"fmt"
	"github.com/Archetarcher/gophkeeper/internal/auth/domain/user"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type SignUp struct {
	Login     string
	Password  string
	Firstname string
	Lastname  string
}

type SignUpHandler struct {
	repo user.Repository
}

func NewSignUpHandler(repo user.Repository) SignUpHandler {
	return SignUpHandler{repo: repo}
}

func (h SignUpHandler) Handle(ctx context.Context, cmd SignUp) error {
	u, err := h.repo.GetByLogin(ctx, cmd.Login)
	if u != nil {
		fmt.Println("starting return error")
		return user.ErrUserAlreadyRegistered
	}
	hash, err := getPasswordHash(cmd.Password)
	if err != nil {
		return errors.Wrap(err, "failed to create password hash")
	}
	newUser, err := user.NewUser(cmd.Login, hash, cmd.Firstname, cmd.Lastname)
	if err != nil {
		return errors.Wrap(err, "auth validation failed")
	}

	return h.repo.Add(ctx, newUser)
}

func getPasswordHash(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
