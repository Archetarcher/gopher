package app

import (
	"github.com/Archetarcher/gophkeeper/internal/vault/app/command"
	"github.com/Archetarcher/gophkeeper/internal/vault/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	RememberCipherLoginData        command.RememberCipherLoginDataHandler
	RememberCipherCustomData       command.RememberCipherCustomDataHandler
	RememberCipherCustomBinaryData command.RememberCipherCustomBinaryDataHandler
	RememberCipherCardData         command.RememberCipherCardDataHandler
}

type Queries struct {
	ShowUserSecrets query.ShowUserSecretsHandler
	ShowSecret      query.ShowSecretHandler
}
