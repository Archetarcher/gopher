package query

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID
	Firstname string
	Lastname  string
	Login     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Token struct {
	Token     string
	ExpiresAt string
}
