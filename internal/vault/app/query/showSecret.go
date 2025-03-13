package query

import (
	"context"
	"github.com/Archetarcher/gophkeeper/internal/common/encryption"
	"github.com/Archetarcher/gophkeeper/internal/vault/domain/secret"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type ShowSecret struct {
	Key    string
	UserId uuid.UUID
}
type ShowSecretHandler struct {
	repo secret.Repository
	enc  *encryption.Asymmetric
}

func NewShowSecretHandler(repo secret.Repository, enc *encryption.Asymmetric) ShowSecretHandler {
	return ShowSecretHandler{repo: repo, enc: enc}
}

func (h ShowSecretHandler) Handle(ctx context.Context, cmd ShowSecret) (*Secret, error) {
	s, err := h.repo.GetSecretByUserAndKey(ctx, cmd.UserId, cmd.Key)
	if err != nil {
		return nil, errors.Wrap(err, "query failed: no entry for given request")
	}

	decKey, err := h.enc.Decrypt(s.GetKey())
	if err != nil {
		return nil, errors.Wrap(err, "query failed: error while key decryption")
	}
	decData, err := h.enc.Decrypt(s.GetData())
	if err != nil {
		return nil, errors.Wrap(err, "query failed: error while data decryption")
	}
	return &Secret{
		ID:         s.GetId(),
		Key:        string(decKey),
		Data:       string(decData),
		SecretType: s.GetType(),
		CreatedAt:  s.GetCreatedAt().String(),
	}, nil

}
