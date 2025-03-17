// Package gophkeeper implements the cobra command for gophkeeper
package gophkeeper

import (
	"context"
	"github.com/Archetarcher/gophkeeper/internal/client/cli/gophkeeper/commands"
	"github.com/spf13/cobra"
)

// NewGophkeeperCommand creates gophkeeper cobra command
func NewGophkeeperCommand(ctx context.Context, server commands.CLIServer) *cobra.Command {
	gophKeeperCmd := &cobra.Command{
		Use:   "gophkeeper",
		Short: "manages your secrets",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
		DisableFlagParsing:    true,
		DisableFlagsInUseLine: true,
	}

	signInCmd := server.NewSignInCommand(ctx)
	signUpCmd := server.NewSignUpCommand(ctx)
	rememberLoginCmd := server.NewRememberLoginCommand(ctx)
	rememberCustomCmd := server.NewRememberCustomCommand(ctx)
	rememberCustomBinaryCmd := server.NewRememberCustomBinaryCommand(ctx)
	rememberCardCmd := server.NewRememberCardCommand(ctx)

	gophKeeperCmd.AddCommand(rememberLoginCmd)
	gophKeeperCmd.AddCommand(rememberCustomCmd)
	gophKeeperCmd.AddCommand(rememberCustomBinaryCmd)
	gophKeeperCmd.AddCommand(rememberCardCmd)
	gophKeeperCmd.AddCommand(signUpCmd)
	gophKeeperCmd.AddCommand(signInCmd)

	return gophKeeperCmd
}
