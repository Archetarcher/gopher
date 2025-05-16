// Package clis implements the command line interface that is used from the entrypoint.
package cli

import (
	"context"
	"github.com/Archetarcher/gophkeeper/internal/client/cli/gophkeeper"
	"github.com/Archetarcher/gophkeeper/internal/client/cli/gophkeeper/commands"
	"github.com/Archetarcher/gophkeeper/internal/client/cli/tui"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gotea",
	Short: "A gotea CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(ctx context.Context, server commands.CLIServer) {
	rootCmd.AddCommand(gophkeeper.NewGophkeeperCommand(ctx, server))
	rootCmd.AddCommand(tui.NewTUICommand())
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
