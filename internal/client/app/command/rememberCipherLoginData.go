package command

import (
	"context"
	"github.com/Archetarcher/gophkeeper/internal/client/provider/vault"
	"github.com/pkg/errors"
)

type RememberCipherLoginData struct {
	Login    string
	Password string
	Uri      string
	Meta     string
}

type RememberCipherLoginDataHandler struct {
	vaultPrv vault.Provider
}

func NewRememberCipherLoginDataHandler(vaultPrv vault.Provider) RememberCipherLoginDataHandler {
	return RememberCipherLoginDataHandler{vaultPrv: vaultPrv}
}

func (h RememberCipherLoginDataHandler) Handle(ctx context.Context, cmd RememberCipherLoginData) error {
	newCipher, err := vault.NewRememberCipherLoginData(cmd.Uri, cmd.Login, cmd.Password, cmd.Meta)
	if err != nil {
		return errors.Wrap(err, "fields validation failed")
	}
	return h.vaultPrv.RememberCipherLogin(ctx, newCipher)
}
