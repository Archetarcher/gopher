package command

import (
	"context"
	"github.com/Archetarcher/gophkeeper/internal/client/provider/auth"
	"github.com/pkg/errors"
)

type SignUp struct {
	Login     string
	Password  string
	Firstname string
	Lastname  string
}

type SignUpHandler struct {
	prv auth.Provider
}

func NewSignUpHandler(prv auth.Provider) SignUpHandler {
	return SignUpHandler{prv: prv}
}

func (h SignUpHandler) Handle(ctx context.Context, cmd SignUp) error {
	newSignUp, err := auth.NewSignUp(cmd.Login, cmd.Password, cmd.Firstname, cmd.Lastname)
	if err != nil {
		return errors.Wrap(err, "sign up validation failed")
	}

	return h.prv.SignUp(ctx, newSignUp)
}
