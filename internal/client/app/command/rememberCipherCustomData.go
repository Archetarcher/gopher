package command

import (
	"context"
	"github.com/Archetarcher/gophkeeper/internal/client/provider/vault"
	"github.com/pkg/errors"
)

type RememberCipherCustomData struct {
	Key   string
	Value string
	Meta  string
}

type RememberCipherCustomDataHandler struct {
	vaultPrv vault.Provider
}

func NewRememberCipherCustomDataHandler(vaultPrv vault.Provider) RememberCipherCustomDataHandler {
	return RememberCipherCustomDataHandler{vaultPrv: vaultPrv}
}

func (h RememberCipherCustomDataHandler) Handle(ctx context.Context, cmd RememberCipherCustomData) error {
	newCipher, err := vault.NewRememberCipherCustomData(cmd.Key, cmd.Value, cmd.Meta)
	if err != nil {
		return errors.Wrap(err, "fields validation failed")
	}
	return h.vaultPrv.RememberCipherCustom(ctx, newCipher)
}
