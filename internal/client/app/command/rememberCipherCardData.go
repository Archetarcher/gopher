package command

import (
	"context"
	"github.com/Archetarcher/gophkeeper/internal/client/provider/vault"
	"github.com/pkg/errors"
)

type RememberCipherCardData struct {
	CardHolderName string
	Brand          string
	Number         string
	ExpMonth       string
	ExpYear        string
	Code           string
	Meta           string
}

type RememberCipherCardDataHandler struct {
	vaultPrv vault.Provider
}

func NewRememberCipherCardDataHandler(vaultPrv vault.Provider) RememberCipherCardDataHandler {
	return RememberCipherCardDataHandler{vaultPrv: vaultPrv}
}

func (h RememberCipherCardDataHandler) Handle(ctx context.Context, cmd RememberCipherCardData) error {
	newCipher, err := vault.NewRememberCipherCardData(cmd.CardHolderName, cmd.Brand, cmd.Number, cmd.ExpMonth, cmd.ExpYear, cmd.Code, cmd.Meta)
	if err != nil {
		return errors.Wrap(err, "fields validation failed")
	}
	return h.vaultPrv.RememberCipherCard(ctx, newCipher)
}
