package command

import (
	"context"
	"github.com/Archetarcher/gophkeeper/internal/common/encryption"
	"github.com/Archetarcher/gophkeeper/internal/vault/domain/secret"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type RememberSecret struct {
	Key        string
	Data       string
	CipherType string
	UserId     uuid.UUID
}

type RememberSecretHandler struct {
	repo secret.Repository
	enc  *encryption.Asymmetric
}

func NewRememberSecretHandler(repo secret.Repository, enc *encryption.Asymmetric) RememberSecretHandler {
	return RememberSecretHandler{repo: repo, enc: enc}
}

func (h RememberSecretHandler) Handle(ctx context.Context, cmd RememberSecret) error {

	encKey, err := h.enc.Encrypt([]byte(cmd.Key))
	if err != nil {
		return errors.Wrap(err, "fields ciphering failed")
	}
	encData, err := h.enc.Encrypt([]byte(cmd.Data))
	if err != nil {
		return errors.Wrap(err, "fields ciphering failed")
	}

	newCipher, err := secret.NewSecret(encData, encKey, cmd.CipherType, cmd.UserId)
	if err != nil {
		return errors.Wrap(err, "fields validation failed")
	}

	return h.repo.Add(ctx, newCipher)
}
