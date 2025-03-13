package app

import (
	"github.com/Archetarcher/gophkeeper/internal/client/app/command"
	"github.com/Archetarcher/gophkeeper/internal/client/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	SignUp                         command.SignUpHandler
	RememberCipherLoginData        command.RememberCipherLoginDataHandler
	RememberCipherCustomData       command.RememberCipherCustomDataHandler
	RememberCipherCustomBinaryData command.RememberCipherCustomBinaryDataHandler
	RememberCipherCardData         command.RememberCipherCardDataHandler
}

type Queries struct {
	SignIn query.SignInHandler
}
