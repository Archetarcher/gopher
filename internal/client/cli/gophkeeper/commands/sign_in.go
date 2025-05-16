package commands

import (
	"context"
	"fmt"
	"github.com/Archetarcher/gophkeeper/internal/client/app/query"
	"github.com/spf13/cobra"
	"log"
)

type SignIn struct {
	Login    string
	Password string
}

var (
	signIn = SignIn{
		Login:    "",
		Password: "",
	}
)

func (c *CLIServer) NewSignInCommand(ctx context.Context) *cobra.Command {
	createCmd := &cobra.Command{
		Use:   "sign-in",
		Short: "sign in user",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("SIGN IN  STARTED")
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := c.app.Queries.SignIn.Handle(ctx, query.SignIn{
				Login:    signIn.Login,
				Password: signIn.Password,
			})
			fmt.Println("err resp")
			fmt.Println(err)
			if err != nil {
				return err
			}
			fmt.Println("USER SUCCESSFULLY SIGNED IN!")
			return nil
		},
	}

	setupLoginFlag()(createCmd)
	setupPasswordFlag()(createCmd)
	return createCmd
}
func setupLoginFlag() commandOpt {
	const flagName = "login"
	return func(cmd *cobra.Command) {
		cmd.Flags().StringVarP(&signIn.Login, flagName, "l", "", "login")
		if err := cmd.MarkFlagRequired(flagName); err != nil {
			log.Fatalf("failed to mark flag as required: %s", err)
		}
	}
}
func setupPasswordFlag() commandOpt {
	const flagName = "password"
	return func(cmd *cobra.Command) {
		cmd.Flags().StringVarP(&signIn.Password, flagName, "p", "", "password")
		if err := cmd.MarkFlagRequired(flagName); err != nil {
			log.Fatalf("failed to mark flag as required: %s", err)
		}
	}
}
