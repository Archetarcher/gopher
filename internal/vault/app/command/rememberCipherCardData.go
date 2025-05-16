package command

import (
	"context"
	"github.com/Archetarcher/gophkeeper/internal/common/encryption"
	cipher "github.com/Archetarcher/gophkeeper/internal/vault/domain/cipherCardData"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type RememberCipherCardData struct {
	CardHolderName string
	Brand          string
	Number         string
	ExpMonth       string
	ExpYear        string
	Code           string
	Meta           *string

	UserId uuid.UUID
}

type RememberCipherCardDataHandler struct {
	repo cipher.Repository
	enc  *encryption.Asymmetric
}

func NewRememberCipherCardDataHandler(repo cipher.Repository, enc *encryption.Asymmetric) RememberCipherCardDataHandler {
	return RememberCipherCardDataHandler{repo: repo, enc: enc}
}

func (h RememberCipherCardDataHandler) Handle(ctx context.Context, cmd RememberCipherCardData) error {
	encCardHolderName, err := h.enc.Encrypt([]byte(cmd.CardHolderName))
	if err != nil {
		return errors.Wrap(err, "fields ciphering failed")
	}

	encBrand, err := h.enc.Encrypt([]byte(cmd.Brand))
	if err != nil {
		return errors.Wrap(err, "fields ciphering failed")
	}

	encNumber, err := h.enc.Encrypt([]byte(cmd.Number))
	if err != nil {
		return errors.Wrap(err, "fields ciphering failed")
	}

	encExpMonth, err := h.enc.Encrypt([]byte(cmd.ExpMonth))
	if err != nil {
		return errors.Wrap(err, "fields ciphering failed")
	}

	encExpYear, err := h.enc.Encrypt([]byte(cmd.ExpYear))
	if err != nil {
		return errors.Wrap(err, "fields ciphering failed")
	}

	encCode, err := h.enc.Encrypt([]byte(cmd.Code))
	if err != nil {
		return errors.Wrap(err, "fields ciphering failed")
	}
	var encMeta []byte
	if cmd.Meta != nil {
		encMeta, err = h.enc.Encrypt([]byte(*cmd.Meta))
		if err != nil {
			return errors.Wrap(err, "fields ciphering failed")
		}
	}

	newCipher, err := cipher.NewCipherCardData(encCardHolderName, encBrand, encNumber, encExpMonth, encExpYear, encCode, encMeta, cmd.UserId)
	if err != nil {
		return errors.Wrap(err, "fields validation failed")
	}

	return h.repo.Add(ctx, newCipher)
}
