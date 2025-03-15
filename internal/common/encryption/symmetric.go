package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// SymmetricEncryption is interface that defines symmetric encryption.
type SymmetricEncryption interface {
	Decrypt(text []byte) ([]byte, error)
	Encrypt(text []byte) ([]byte, error)
}

// Symmetric is struct for symmetric encryption.
type Symmetric struct {
	key string
}

// NewSymmetric creates instance of Symmetric, use key for encryption.
func NewSymmetric(key string) *Symmetric {
	return &Symmetric{key: key}
}

// Decrypt is a function that uses symmetric decryption to the given slice of bytes
func (s *Symmetric) Decrypt(ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(s.key))
	if err != nil {

		logrus.Error("error symmetric decryption ", err)
		return nil, errors.Wrap(err, "error decryption")
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		logrus.Error("error symmetric decryption ", err)
		return nil, errors.Wrap(err, "error decryption")
	}

	plaintext, err := gcm.Open(nil, ciphertext[len(ciphertext)-12:], ciphertext[:len(ciphertext)-12], nil)
	if err != nil {
		logrus.Error("error symmetric decryption ", err)
		return nil, errors.Wrap(err, "error decryption")
	}

	return plaintext, nil
}
func (s *Symmetric) Encrypt(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(s.key))
	if err != nil {
		logrus.Error("error symmetric encryption ", err)
		return nil, errors.Wrap(err, "error encryption")
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		logrus.Error("error symmetric encryption ", err)
		return nil, errors.Wrap(err, "error encryption")
	}

	nonce := make([]byte, 12)
	if _, err := rand.Read(nonce); err != nil {
		logrus.Error("error symmetric encryption ", err)
		return nil, errors.Wrap(err, "error encryption")
	}

	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)
	ciphertext = append(ciphertext, nonce...)
	return ciphertext, nil
}
