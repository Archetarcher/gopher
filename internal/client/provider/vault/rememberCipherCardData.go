package vault

import (
	"github.com/pkg/errors"
)

var (
	ErrInvalidCipherCardData = errors.New("cipher has to have  valid fields")
)

// RememberCipherCardData is an aggregate for auth
type RememberCipherCardData struct {
	cardHolderName string
	brand          string
	number         string
	expMonth       string
	expYear        string
	code           string
	meta           string
}

// NewRememberCipherCardData is a Factory to create a new CipherCardData aggregate
// It will validate that the data, key, userId, cipherType are not empty
func NewRememberCipherCardData(cardHolderName, brand, number, expMonth, expYear, code, meta string) (*RememberCipherCardData, error) {
	if cardHolderName == "" {
		return nil, errors.Wrap(ErrInvalidCipherCardData, "cardHolderName does not provided")
	}
	if brand == "" {
		return nil, errors.Wrap(ErrInvalidCipherCardData, "brand does not provided")
	}
	if number == "" {
		return nil, errors.Wrap(ErrInvalidCipherCardData, "number does not provided")
	}
	if expMonth == "" {
		return nil, errors.Wrap(ErrInvalidCipherCardData, "expMonth does not provided")
	}
	if expYear == "" {
		return nil, errors.Wrap(ErrInvalidCipherCardData, "expYear does not provided")
	}
	if code == "" {
		return nil, errors.Wrap(ErrInvalidCipherCardData, "code does not provided")
	}

	return &RememberCipherCardData{
		cardHolderName: cardHolderName,
		brand:          brand,
		number:         number,
		expMonth:       expMonth,
		expYear:        expYear,
		code:           code,
		meta:           meta,
	}, nil
}
