package command

import (
	"context"
	"github.com/Archetarcher/gophkeeper/internal/common/encryption"
	cipher "github.com/Archetarcher/gophkeeper/internal/vault/domain/cipherCustomBinaryData"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type RememberCipherCustomBinaryData struct {
	Key   string
	Value string
	Meta  *string

	UserId uuid.UUID
}

type RememberCipherCustomBinaryDataHandler struct {
	repo cipher.Repository
	enc  *encryption.Asymmetric
}

func NewRememberCipherCustomBinaryDataHandler(repo cipher.Repository, enc *encryption.Asymmetric) RememberCipherCustomBinaryDataHandler {
	return RememberCipherCustomBinaryDataHandler{repo: repo, enc: enc}
}

func (h RememberCipherCustomBinaryDataHandler) Handle(ctx context.Context, cmd RememberCipherCustomBinaryData) error {
	encKey, err := h.enc.Encrypt([]byte(cmd.Key))
	if err != nil {
		return errors.Wrap(err, "fields ciphering failed")
	}
	encValue, err := h.enc.Encrypt([]byte(cmd.Value))
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

	newCipher, err := cipher.NewCipherCustomBinaryData(encKey, encValue, encMeta, cmd.UserId)
	if err != nil {
		return errors.Wrap(err, "fields validation failed")
	}

	return h.repo.Add(ctx, newCipher)
}
