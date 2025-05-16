package secret

import (
	"github.com/pkg/errors"
)

var (
	ErrSecretTypeIncorrectInput = errors.New("invalid secret type provided")
)

type Type struct {
	s string
}

var (
	Auth         = Type{"auth"}
	Custom       = Type{"custom"}
	CustomBinary = Type{"custom_binary"}
	Card         = Type{"card"}
)

func NewSecretTypeFromString(userType string) (Type, error) {
	switch userType {
	case "auth":
		return Auth, nil
	case "custom":
		return Custom, nil
	case "custom_binary":
		return CustomBinary, nil
	case "card":
		return Card, nil
	}

	return Type{}, ErrSecretTypeIncorrectInput
}
