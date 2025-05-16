package query

import (
	"context"
	"github.com/Archetarcher/gophkeeper/internal/auth/domain/user"
	"github.com/Archetarcher/gophkeeper/internal/common/auth"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type SignIn struct {
	Login    string
	Password string
}
type SignInHandler struct {
	repo     user.Repository
	tokenCfg auth.JWTTokenConfig
}

func NewSignInHandler(repo user.Repository, token auth.JWTTokenConfig) SignInHandler {
	return SignInHandler{repo: repo, tokenCfg: token}
}

func (h SignInHandler) Handle(ctx context.Context, cmd SignIn) (*Token, error) {
	u, err := h.repo.GetByLogin(ctx, cmd.Login)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, user.ErrUserNotFound
	}

	if bcrypt.CompareHashAndPassword([]byte(u.GetHash()), []byte(cmd.Password)) != nil {
		return nil, errors.Wrap(user.ErrSignIn, "bad credentials, login password pare are not valid")
	}

	token, err := h.tokenCfg.CreateToken(u.GetId())
	if err != nil {
		return nil, errors.Wrap(err, "failed to create token")
	}
	return &Token{
		Token:     token,
		ExpiresAt: h.tokenCfg.GetTokenExpiration(),
	}, nil
}
