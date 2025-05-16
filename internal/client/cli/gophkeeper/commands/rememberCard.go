package commands

import (
	"context"
	"fmt"
	"github.com/Archetarcher/gophkeeper/internal/client/app/command"
	"github.com/spf13/cobra"
	"log"
)

type RememberCard struct {
	Brand           string
	CardHolderName  string
	Code            string
	ExpirationMonth string
	ExpirationYear  string
	Number          string
}

var (
	rememberCard = RememberCard{
		Brand:           "",
		CardHolderName:  "",
		Code:            "",
		ExpirationMonth: "",
		ExpirationYear:  "",
		Number:          "",
	}
)

func (c *CLIServer) NewRememberCardCommand(ctx context.Context) *cobra.Command {
	createCmd := &cobra.Command{
		Use:   "remember-card",
		Short: "remembers card secret data",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("REMEMBER CARD DATA  STARTED")
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			err := c.app.Commands.RememberCipherCardData.Handle(ctx, command.RememberCipherCardData{
				Brand:          rememberCard.Brand,
				CardHolderName: rememberCard.CardHolderName,
				Code:           rememberCard.Code,
				ExpYear:        rememberCard.ExpirationYear,
				ExpMonth:       rememberCard.ExpirationMonth,
				Number:         rememberCard.Number,
			})
			if err != nil {
				return err
			}
			fmt.Println("USER DATA SUCCESSFULLY REMEMBERED!")
			return nil
		},
	}

	setupRememberBrandFlag()(createCmd)
	setupRememberNameFlag()(createCmd)
	setupRememberCodeFlag()(createCmd)
	setupRememberMonthFlag()(createCmd)
	setupRememberYearFlag()(createCmd)
	setupRememberNumberFlag()(createCmd)
	return createCmd
}
func setupRememberBrandFlag() commandOpt {
	const flagName = "brand"
	return func(cmd *cobra.Command) {
		cmd.Flags().StringVarP(&rememberCard.Brand, flagName, "b", "", "brand")
		if err := cmd.MarkFlagRequired(flagName); err != nil {
			log.Fatalf("failed to mark flag as required: %s", err)
		}
	}
}
func setupRememberNameFlag() commandOpt {
	const flagName = "card golder name"
	return func(cmd *cobra.Command) {
		cmd.Flags().StringVarP(&rememberCard.CardHolderName, flagName, "h", "", "card holder name")
		if err := cmd.MarkFlagRequired(flagName); err != nil {
			log.Fatalf("failed to mark flag as required: %s", err)
		}
	}
}
func setupRememberCodeFlag() commandOpt {
	const flagName = "code"
	return func(cmd *cobra.Command) {
		cmd.Flags().StringVarP(&rememberCard.Code, flagName, "c", "", "code")
		if err := cmd.MarkFlagRequired(flagName); err != nil {
			log.Fatalf("failed to mark flag as required: %s", err)
		}
	}
}
func setupRememberMonthFlag() commandOpt {
	const flagName = "expiration month"
	return func(cmd *cobra.Command) {
		cmd.Flags().StringVarP(&rememberCard.ExpirationMonth, flagName, "m", "", "expiration month")
		if err := cmd.MarkFlagRequired(flagName); err != nil {
			log.Fatalf("failed to mark flag as required: %s", err)
		}
	}
}
func setupRememberYearFlag() commandOpt {
	const flagName = "expiration year"
	return func(cmd *cobra.Command) {
		cmd.Flags().StringVarP(&rememberCard.ExpirationYear, flagName, "y", "", "expiration year")
		if err := cmd.MarkFlagRequired(flagName); err != nil {
			log.Fatalf("failed to mark flag as required: %s", err)
		}
	}
}
func setupRememberNumberFlag() commandOpt {
	const flagName = "number"
	return func(cmd *cobra.Command) {
		cmd.Flags().StringVarP(&rememberCard.Number, flagName, "n", "", "number")
		if err := cmd.MarkFlagRequired(flagName); err != nil {
			log.Fatalf("failed to mark flag as required: %s", err)
		}
	}
}
