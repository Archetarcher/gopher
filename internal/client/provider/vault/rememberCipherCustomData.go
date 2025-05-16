package vault

import (
	"github.com/pkg/errors"
)

var (
	ErrInvalidCipherCustomData = errors.New("cipher has to have  valid fields")
)

// RememberCipherCustomData is an aggregate for auth
type RememberCipherCustomData struct {
	key   string
	value string
	meta  string
}

// NewRememberCipherCustomData is a Factory to create a new CipherCustomData aggregate
// It will validate that the data, key, userId, cipherType are not empty
func NewRememberCipherCustomData(key, value, meta string) (*RememberCipherCustomData, error) {
	if key == "" {
		return nil, errors.Wrap(ErrInvalidCipherCustomBinaryData, "key does not provided")
	}
	if value == "" {
		return nil, errors.Wrap(ErrInvalidCipherCustomBinaryData, "value does not provided")
	}

	return &RememberCipherCustomData{
		key:   key,
		value: value,
		meta:  meta,
	}, nil
}

func (c *RememberCipherCustomData) GetKey() string {
	return c.key
}

func (c *RememberCipherCustomData) GetValue() string {
	return c.value
}

func (c *RememberCipherCustomData) GetMeta() string {
	return c.meta
}
