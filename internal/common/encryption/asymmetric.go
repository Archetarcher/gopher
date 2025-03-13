package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
)

// AsymmetricEncryption is an interface that defines asymmetric encryption.
type AsymmetricEncryption interface {
	Decrypt(text []byte) ([]byte, error)
	Encrypt(text []byte) ([]byte, error)
}

// Asymmetric is a struct for asymmetric encryption.
type Asymmetric struct {
	publicKeyPath  string
	privateKeyPath string
}

// NewAsymmetric creates instance of Asymmetric.
func NewAsymmetric(publicKeyPath string, privateKeyPath string) *Asymmetric {
	return &Asymmetric{publicKeyPath: publicKeyPath, privateKeyPath: privateKeyPath}
}

// Decrypt is a function that uses asymmetric decryption to the given slice of bytes
func (a *Asymmetric) Decrypt(ciphertext []byte) ([]byte, error) {
	privateKeyPEM, err := os.ReadFile(a.privateKeyPath)
	if err != nil {
		logrus.Error("failed to read key file", err)
		return nil, errors.Wrap(err, "failed to read key file")
	}
	privateKeyBlock, _ := pem.Decode(privateKeyPEM)
	privateKey, err := x509.ParsePKCS8PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		logrus.Error("error load key", err)
		return nil, errors.Wrap(err, "error load key")
	}
	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), ciphertext)
	if err != nil {
		logrus.Error("error decrypt", err)
		return nil, errors.Wrap(err, "error decrypt")
	}

	return decrypted, nil
}

// Encrypt is a function that uses asymmetric encryption to the given slice of bytes
func (a *Asymmetric) Encrypt(js []byte) ([]byte, error) {
	publicKeyPEM, err := os.ReadFile(a.publicKeyPath)
	if err != nil {
		logrus.Error("error asymmetric encryption", err)
		return nil, errors.Wrap(err, "error asymmetric encryption")

	}
	publicKeyBlock, _ := pem.Decode(publicKeyPEM)

	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		logrus.Error("error asymmetric encryption", err)
		return nil, errors.Wrap(err, "error asymmetric encryption")
	}

	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey.(*rsa.PublicKey), js)
	if err != nil {
		logrus.Error("error asymmetric encryption", err)
		return nil, errors.Wrap(err, "error asymmetric encryption")
	}
	return ciphertext, nil
}

// GenKey generates crypto key
func GenKey(n int) ([]byte, error) {
	rnd := make([]byte, n)

	nrnd, err := rand.Read(rnd)
	if err != nil {
		return nil, err
	} else if nrnd != n {
		return nil, fmt.Errorf(`nrnd %d != n %d`, nrnd, n)
	}
	for i := range rnd {
		rnd[i] = 'A' + rnd[i]%26
	}
	return rnd, nil
}
