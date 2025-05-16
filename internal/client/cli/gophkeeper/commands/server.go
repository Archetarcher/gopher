package commands

import (
	"github.com/Archetarcher/gophkeeper/internal/client/app"
	"github.com/spf13/cobra"
)

type CLIServer struct {
	app app.Application
}

func NewCLI(app app.Application) CLIServer {
	return CLIServer{app: app}
}

type commandOpt func(command *cobra.Command)
