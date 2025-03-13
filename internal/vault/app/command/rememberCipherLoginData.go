package command

import (
	"context"
	"github.com/Archetarcher/gophkeeper/internal/common/encryption"
	cipher "github.com/Archetarcher/gophkeeper/internal/vault/domain/cipherLoginData"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type RememberCipherLoginData struct {
	Uri      string
	Login    string
	Password string
	Meta     *string

	UserId uuid.UUID
}

type RememberCipherLoginDataHandler struct {
	repo cipher.Repository
	enc  *encryption.Asymmetric
}

func NewRememberCipherLoginDataHandler(repo cipher.Repository, enc *encryption.Asymmetric) RememberCipherLoginDataHandler {
	return RememberCipherLoginDataHandler{repo: repo, enc: enc}
}

func (h RememberCipherLoginDataHandler) Handle(ctx context.Context, cmd RememberCipherLoginData) error {

	encLogin, err := h.enc.Encrypt([]byte(cmd.Login))
	if err != nil {
		return errors.Wrap(err, "fields ciphering failed")
	}
	encPassword, err := h.enc.Encrypt([]byte(cmd.Password))
	if err != nil {
		return errors.Wrap(err, "fields ciphering failed")
	}
	encUri, err := h.enc.Encrypt([]byte(cmd.Uri))
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
	newCipher, err := cipher.NewCipherLoginData(encUri, encLogin, encPassword, encMeta, cmd.UserId)
	if err != nil {
		return errors.Wrap(err, "fields validation failed")
	}

	return h.repo.Add(ctx, newCipher)
}
