package vault

import (
	"github.com/pkg/errors"
)

var (
	ErrInvalidCipherCustomBinaryData = errors.New("cipher has to have  valid fields")
)

// RememberCipherCustomBinaryData is an aggregate for auth
type RememberCipherCustomBinaryData struct {
	key   string
	value string
	meta  string
}

// NewRememberCipherCustomBinaryData is a Factory to create a new CipherCustomBinaryData aggregate
// It will validate that the data, key, userId, cipherType are not empty
func NewRememberCipherCustomBinaryData(key, value, meta string) (*RememberCipherCustomBinaryData, error) {
	if key == "" {
		return nil, errors.Wrap(ErrInvalidCipherCustomBinaryData, "key does not provided")
	}
	if value == "" {
		return nil, errors.Wrap(ErrInvalidCipherCustomBinaryData, "value does not provided")
	}
	return &RememberCipherCustomBinaryData{
		key:   key,
		value: value,
		meta:  meta,
	}, nil
}
func (c *RememberCipherCustomBinaryData) GetKey() string {
	return c.key
}

func (c *RememberCipherCustomBinaryData) GetValue() string {
	return c.value
}

func (c *RememberCipherCustomBinaryData) GetMeta() string {
	return c.meta
}
