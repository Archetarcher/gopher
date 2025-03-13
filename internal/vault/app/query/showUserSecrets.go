package query

import (
	"context"
	"github.com/Archetarcher/gophkeeper/internal/common/encryption"
	"github.com/Archetarcher/gophkeeper/internal/vault/domain/secret"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type ShowUserSecrets struct {
	UserId uuid.UUID
}
type ShowUserSecretsHandler struct {
	repo secret.Repository
	enc  *encryption.Asymmetric
}

func NewShowUserSecretsHandler(repo secret.Repository, enc *encryption.Asymmetric) ShowUserSecretsHandler {
	return ShowUserSecretsHandler{repo: repo, enc: enc}
}

func (h ShowUserSecretsHandler) Handle(ctx context.Context, cmd ShowUserSecrets) ([]Secret, error) {
	var secrets []Secret
	rc, err := h.repo.GetAllSecretByUser(ctx, cmd.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch existing secrets")
	}
	for _, c := range rc {
		decKey, err := h.enc.Decrypt([]byte(c.GetKey()))
		if err != nil {
			return nil, errors.Wrap(err, "failed to decrypt secret key")
		}
		decData, err := h.enc.Decrypt([]byte(c.GetData()))
		if err != nil {
			return nil, errors.Wrap(err, "failed to decrypt secret data")
		}
		secrets = append(secrets, Secret{
			ID:         c.GetId(),
			Key:        string(decKey),
			Data:       string(decData),
			SecretType: c.GetType(),
			CreatedAt:  c.GetCreatedAt().String(),
		})
	}
	return secrets, nil
}
