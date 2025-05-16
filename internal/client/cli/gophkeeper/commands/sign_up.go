package commands

import (
	"context"
	"fmt"
	"github.com/Archetarcher/gophkeeper/internal/client/app/command"
	"github.com/spf13/cobra"
	"log"
)

type SignUp struct {
	Firstname string
	Lastname  string
	Login     string
	Password  string
}

var (
	signUp = SignUp{
		Login:    "",
		Password: "",
	}
)

func (c *CLIServer) NewSignUpCommand(ctx context.Context) *cobra.Command {
	createCmd := &cobra.Command{
		Use:   "sign-up",
		Short: "sign up user",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("SIGN UP  STARTED")
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			err := c.app.Commands.SignUp.Handle(ctx, command.SignUp{
				Login:     signUp.Login,
				Password:  signUp.Password,
				Firstname: signUp.Firstname,
				Lastname:  signUp.Lastname,
			})
			if err != nil {
				return err
			}
			fmt.Println("USER SUCCESSFULLY SIGNED UP!")
			return nil
		},
	}

	setupSignUpLoginFlag()(createCmd)
	setupSignUpPasswordFlag()(createCmd)
	setupFirstnameFlag()(createCmd)
	setupLastnameFlag()(createCmd)
	return createCmd
}
func setupSignUpLoginFlag() commandOpt {
	const flagName = "login"
	return func(cmd *cobra.Command) {
		cmd.Flags().StringVarP(&signUp.Login, flagName, "l", "", "login")
		if err := cmd.MarkFlagRequired(flagName); err != nil {
			log.Fatalf("failed to mark flag as required: %s", err)
		}
	}
}
func setupSignUpPasswordFlag() commandOpt {
	const flagName = "password"
	return func(cmd *cobra.Command) {
		cmd.Flags().StringVarP(&signUp.Password, flagName, "p", "", "password")
		if err := cmd.MarkFlagRequired(flagName); err != nil {
			log.Fatalf("failed to mark flag as required: %s", err)
		}
	}
}

func setupFirstnameFlag() commandOpt {
	const flagName = "firstname"
	return func(cmd *cobra.Command) {
		cmd.Flags().StringVarP(&signUp.Firstname, flagName, "f", "", "firstname")
		if err := cmd.MarkFlagRequired(flagName); err != nil {
			log.Fatalf("failed to mark flag as required: %s", err)
		}
	}
}
func setupLastnameFlag() commandOpt {
	const flagName = "lastname"
	return func(cmd *cobra.Command) {
		cmd.Flags().StringVarP(&signUp.Lastname, flagName, "n", "", "lastname")
		if err := cmd.MarkFlagRequired(flagName); err != nil {
			log.Fatalf("failed to mark flag as required: %s", err)
		}
	}
}
