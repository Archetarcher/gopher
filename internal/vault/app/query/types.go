package query

import (
	"github.com/google/uuid"
)

type Secret struct {
	ID         uuid.UUID
	Key        string
	Data       string
	SecretType string
	CreatedAt  string
}
