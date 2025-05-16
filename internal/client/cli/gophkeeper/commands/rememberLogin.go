package commands

import (
	"context"
	"fmt"
	"github.com/Archetarcher/gophkeeper/internal/client/app/command"
	"github.com/spf13/cobra"
	"log"
)

type RememberLogin struct {
	Login    string
	Password string
	Uri      string
}

var (
	rememberLogin = RememberLogin{
		Login:    "",
		Password: "",
		Uri:      "",
	}
)

func (c *CLIServer) NewRememberLoginCommand(ctx context.Context) *cobra.Command {
	createCmd := &cobra.Command{
		Use:   "remember-login",
		Short: "remembers sites login data",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("REMEMBER LOGIN  STARTED")
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			err := c.app.Commands.RememberCipherLoginData.Handle(ctx, command.RememberCipherLoginData{
				Login:    rememberLogin.Login,
				Password: rememberLogin.Password,
				Uri:      rememberLogin.Uri,
			})
			if err != nil {
				return err
			}
			fmt.Println("USER DATA SUCCESSFULLY REMEMBERED!")
			return nil
		},
	}

	setupRememberLoginFlag()(createCmd)
	setupRememberPasswordFlag()(createCmd)
	setupRememberUriFlag()(createCmd)
	return createCmd
}
func setupRememberLoginFlag() commandOpt {
	const flagName = "login"
	return func(cmd *cobra.Command) {
		cmd.Flags().StringVarP(&rememberLogin.Login, flagName, "l", "", "login")
		if err := cmd.MarkFlagRequired(flagName); err != nil {
			log.Fatalf("failed to mark flag as required: %s", err)
		}
	}
}
func setupRememberPasswordFlag() commandOpt {
	const flagName = "password"
	return func(cmd *cobra.Command) {
		cmd.Flags().StringVarP(&rememberLogin.Password, flagName, "p", "", "password")
		if err := cmd.MarkFlagRequired(flagName); err != nil {
			log.Fatalf("failed to mark flag as required: %s", err)
		}
	}
}
func setupRememberUriFlag() commandOpt {
	const flagName = "uri"
	return func(cmd *cobra.Command) {
		cmd.Flags().StringVarP(&rememberLogin.Uri, flagName, "u", "", "site uri")
		if err := cmd.MarkFlagRequired(flagName); err != nil {
			log.Fatalf("failed to mark flag as required: %s", err)
		}
	}
}
