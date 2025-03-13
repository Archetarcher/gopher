package query

import (
	"context"
	"github.com/Archetarcher/gophkeeper/internal/client/provider/auth"
	"github.com/Archetarcher/gophkeeper/internal/common/provider"
	"github.com/pkg/errors"
)

type SignIn struct {
	Login    string
	Password string
}
type SignInHandler struct {
	prv auth.Provider
}

func NewSignInHandler(prv auth.Provider) SignInHandler {
	return SignInHandler{prv: prv}
}

func (h SignInHandler) Handle(ctx context.Context, cmd SignIn) (*provider.Token, error) {
	newSignIn, err := auth.NewSignIn(cmd.Login, cmd.Password)
	if err != nil {
		return nil, errors.Wrap(err, "sign in validation failed")
	}

	return h.prv.SignIn(ctx, newSignIn)
}
