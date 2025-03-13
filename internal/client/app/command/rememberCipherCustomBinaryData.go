package command

import (
	"context"
	"github.com/Archetarcher/gophkeeper/internal/client/provider/vault"
	"github.com/pkg/errors"
)

type RememberCipherCustomBinaryData struct {
	Key   string
	Value string
	Meta  string
}

type RememberCipherCustomBinaryDataHandler struct {
	vaultPrv vault.Provider
}

func NewRememberCipherCustomBinaryDataHandler(vaultPrv vault.Provider) RememberCipherCustomBinaryDataHandler {
	return RememberCipherCustomBinaryDataHandler{vaultPrv: vaultPrv}
}

func (h RememberCipherCustomBinaryDataHandler) Handle(ctx context.Context, cmd RememberCipherCustomBinaryData) error {
	newCipher, err := vault.NewRememberCipherCustomBinaryData(cmd.Key, cmd.Value, cmd.Meta)
	if err != nil {
		return errors.Wrap(err, "fields validation failed")
	}
	return h.vaultPrv.RememberCipherCustomBinary(ctx, newCipher)
}
