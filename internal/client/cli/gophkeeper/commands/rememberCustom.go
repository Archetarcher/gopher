package commands

import (
	"context"
	"fmt"
	"github.com/Archetarcher/gophkeeper/internal/client/app/command"
	"github.com/spf13/cobra"
	"log"
)

type RememberCustom struct {
	Key   string
	Value string
}

var (
	rememberCustom = RememberCustom{
		Key:   "",
		Value: "",
	}
)

func (c *CLIServer) NewRememberCustomCommand(ctx context.Context) *cobra.Command {
	createCmd := &cobra.Command{
		Use:   "remember-custom",
		Short: "remembers custom secret data",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("REMEMBER CUSTOM DATA  STARTED")
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			err := c.app.Commands.RememberCipherCustomData.Handle(ctx, command.RememberCipherCustomData{
				Key:   rememberCustom.Key,
				Value: rememberCustom.Value,
			})
			if err != nil {
				return err
			}
			fmt.Println("USER DATA SUCCESSFULLY REMEMBERED!")
			return nil
		},
	}

	setupRememberKeyFlag()(createCmd)
	setupRememberValueFlag()(createCmd)
	return createCmd
}
func setupRememberKeyFlag() commandOpt {
	const flagName = "key"
	return func(cmd *cobra.Command) {
		cmd.Flags().StringVarP(&rememberCustom.Key, flagName, "k", "", "key")
		if err := cmd.MarkFlagRequired(flagName); err != nil {
			log.Fatalf("failed to mark flag as required: %s", err)
		}
	}
}
func setupRememberValueFlag() commandOpt {
	const flagName = "value"
	return func(cmd *cobra.Command) {
		cmd.Flags().StringVarP(&rememberCustom.Value, flagName, "v", "", "value")
		if err := cmd.MarkFlagRequired(flagName); err != nil {
			log.Fatalf("failed to mark flag as required: %s", err)
		}
	}
}
