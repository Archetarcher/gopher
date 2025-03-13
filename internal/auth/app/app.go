package app

import (
	"github.com/Archetarcher/gophkeeper/internal/auth/app/command"
	"github.com/Archetarcher/gophkeeper/internal/auth/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	SignUp command.SignUpHandler
}

type Queries struct {
	SignIn query.SignInHandler
}
