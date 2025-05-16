package vault

import (
	"github.com/google/uuid"
	"time"
)

// Cipher is an aggregate for auth
type Cipher struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	MetaData  []byte
	DeletedAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
