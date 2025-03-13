package http

import (
	"context"
	b64 "encoding/base64"
	"fmt"
	"github.com/Archetarcher/gophkeeper/internal/client/provider/vault"
	"github.com/Archetarcher/gophkeeper/internal/common/encryption"
	"github.com/Archetarcher/gophkeeper/internal/common/provider"
	"github.com/pkg/errors"
	"net/http"
	"sync"
)

type Provider struct {
	config *provider.Config

	sync.Mutex
}

func New(config *provider.Config, addr string) *Provider {
	config.RunAddr += addr

	return &Provider{
		config: config,
	}
}

func (r *Provider) RememberCipherLogin(ctx context.Context, u *vault.RememberCipherLoginData) error {
	r.Lock()
	defer r.Unlock()

	fmt.Println(r.config.Token)
	if r.config.Token.IsExpired() {
		return vault.ErrTokenExpired
	}
	url := r.config.RunAddr + "/login-data/remember"

	res, err := r.config.Client.
		R().
		SetBody(u).
		Post(url)
	if err != nil {
		return errors.Wrap(err, "provider: could not create request")
	}

	if res.StatusCode() != http.StatusOK {
		return errors.Wrap(err, "provider: responded with error")

	}
	return nil
}
func (r *Provider) RememberCipherCustom(ctx context.Context, u *vault.RememberCipherCustomData) error {
	r.Lock()
	defer r.Unlock()

	fmt.Println(r.config.Token)
	if r.config.Token.IsExpired() {
		return vault.ErrTokenExpired
	}
	url := r.config.RunAddr + "/custom-data/remember"

	res, err := r.config.Client.
		R().
		SetBody(u).
		Post(url)
	if err != nil {
		return errors.Wrap(err, "provider: could not create request")
	}

	if res.StatusCode() != http.StatusOK {
		return errors.Wrap(err, "provider: responded with error")

	}
	return nil
}
func (r *Provider) RememberCipherCustomBinary(ctx context.Context, u *vault.RememberCipherCustomBinaryData) error {
	r.Lock()
	defer r.Unlock()

	fmt.Println(r.config.Token)
	if r.config.Token.IsExpired() {
		return vault.ErrTokenExpired
	}
	url := r.config.RunAddr + "/custom-binary-data/remember"

	res, err := r.config.Client.
		R().
		SetBody(u).
		Post(url)
	if err != nil {
		return errors.Wrap(err, "provider: could not create request")
	}

	if res.StatusCode() != http.StatusOK {
		return errors.Wrap(err, "provider: responded with error")

	}
	return nil
}
func (r *Provider) RememberCipherCard(ctx context.Context, u *vault.RememberCipherCardData) error {
	r.Lock()
	defer r.Unlock()

	fmt.Println(r.config.Token)
	if r.config.Token.IsExpired() {
		return vault.ErrTokenExpired
	}
	url := r.config.RunAddr + "/card-data/remember"

	res, err := r.config.Client.
		R().
		SetBody(u).
		Post(url)
	if err != nil {
		return errors.Wrap(err, "provider: could not create request")
	}

	if res.StatusCode() != http.StatusOK {
		return errors.Wrap(err, "provider: responded with error")

	}
	return nil
}
func (r *Provider) StartSession(ctx context.Context, enc *encryption.Asymmetric) error {
	r.Lock()
	defer r.Unlock()

	url := r.config.RunAddr + "/session"

	key, gErr := encryption.GenKey(16)
	if gErr != nil {
		return errors.Wrap(gErr, "provider: failed to generate crypto key")
	}
	encryptedKey, eErr := enc.Encrypt(key)
	if eErr != nil {
		return eErr
	}

	sEnc := b64.StdEncoding.EncodeToString([]byte(encryptedKey))

	res, err := r.config.Client.
		R().
		SetBody(map[string]string{
			"key": sEnc,
		}).
		Post(url)

	if err != nil {
		return errors.Wrap(gErr, "provider: could not create request")
	}

	if res.StatusCode() != http.StatusOK {
		return errors.Wrap(gErr, "provider: responded with error creating session")
	}

	r.config.Session.Key = string(key)
	return nil
}
