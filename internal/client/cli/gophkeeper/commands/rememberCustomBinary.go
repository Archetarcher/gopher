package commands

import (
	"context"
	"fmt"
	"github.com/Archetarcher/gophkeeper/internal/client/app/command"
	"github.com/spf13/cobra"
	"log"
)

type RememberCustomBinary struct {
	Key   string
	Value string
}

var (
	rememberCustomBinary = RememberCustomBinary{
		Key:   "",
		Value: "",
	}
)

func (c *CLIServer) NewRememberCustomBinaryCommand(ctx context.Context) *cobra.Command {
	createCmd := &cobra.Command{
		Use:   "remember-custom-binary",
		Short: "remembers custom binary secret data",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("REMEMBER CUSTOM DATA  STARTED")
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			err := c.app.Commands.RememberCipherCustomBinaryData.Handle(ctx, command.RememberCipherCustomBinaryData{
				Key:   rememberCustomBinary.Key,
				Value: rememberCustomBinary.Value,
			})
			if err != nil {
				return err
			}
			fmt.Println("USER DATA SUCCESSFULLY REMEMBERED!")
			return nil
		},
	}

	setupRememberBinaryKeyFlag()(createCmd)
	setupRememberBinaryValueFlag()(createCmd)
	return createCmd
}
func setupRememberBinaryKeyFlag() commandOpt {
	const flagName = "key"
	return func(cmd *cobra.Command) {
		cmd.Flags().StringVarP(&rememberCustom.Key, flagName, "k", "", "key")
		if err := cmd.MarkFlagRequired(flagName); err != nil {
			log.Fatalf("failed to mark flag as required: %s", err)
		}
	}
}
func setupRememberBinaryValueFlag() commandOpt {
	const flagName = "value"
	return func(cmd *cobra.Command) {
		cmd.Flags().StringVarP(&rememberCustom.Value, flagName, "v", "", "value")
		if err := cmd.MarkFlagRequired(flagName); err != nil {
			log.Fatalf("failed to mark flag as required: %s", err)
		}
	}
}
