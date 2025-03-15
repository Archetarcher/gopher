package http

import (
	"context"
	"fmt"
	"github.com/Archetarcher/gophkeeper/internal/client/provider/vault"
	"github.com/Archetarcher/gophkeeper/internal/common/provider"
	"github.com/Archetarcher/gophkeeper/internal/common/server/httperr"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"net/http"
	"sync"
)

type Provider struct {
	config  *provider.Config
	runAddr string
	client  *resty.Client

	sync.Mutex
}

func New(config *provider.Config, runAddr string, mls ...provider.MiddlewareFunc) *Provider {
	runAddr += "/api"
	client := resty.New()
	for _, m := range mls {
		client.OnBeforeRequest(resty.RequestMiddleware(m))
	}
	return &Provider{
		config:  config,
		runAddr: runAddr,
		client:  client,
	}
}

func (r *Provider) RememberCipherLogin(ctx context.Context, c *vault.RememberCipherLoginData) error {
	r.Lock()
	defer r.Unlock()

	fmt.Println("r.config.Token")
	fmt.Println(r.config.Token)
	if r.config.Token.IsExpired() {
		return vault.ErrTokenExpired
	}
	url := r.runAddr + "/login-data/remember"

	res, err := r.client.
		R().
		SetHeader("Authorization", "Bearer "+r.config.Token.Token).
		SetBody(RememberCipherLoginDataRequest{
			Login:    c.GetLogin(),
			Meta:     c.GetMeta(),
			Password: c.GetPassword(),
			Uri:      c.GetUri(),
		}).
		Post(url)
	if err != nil {
		return errors.Wrap(err, "provider: could not create request")
	}

	if res.StatusCode() != http.StatusCreated {
		parsedErr := httperr.ParseErrorResponseMessage(res.Body())
		fmt.Println(parsedErr)
		if parsedErr != "" {
			return errors.Wrap(vault.ErrFailedToRememberCipher, parsedErr)
		}
		return errors.Wrap(vault.ErrFailedToRememberCipher, res.String())
	}
	return nil
}
func (r *Provider) RememberCipherCustom(ctx context.Context, c *vault.RememberCipherCustomData) error {
	r.Lock()
	defer r.Unlock()

	fmt.Println(r.config.Token)
	if r.config.Token.IsExpired() {
		return vault.ErrTokenExpired
	}
	url := r.runAddr + "/custom-data/remember"

	res, err := r.client.
		R().
		SetHeader("Authorization", "Bearer "+r.config.Token.Token).
		SetBody(RememberCipherCustomDataRequest{
			Key:   c.GetKey(),
			Meta:  c.GetMeta(),
			Value: c.GetValue(),
		}).
		Post(url)
	if err != nil {
		return errors.Wrap(err, "provider: could not create request")
	}

	if res.StatusCode() != http.StatusCreated {
		parsedErr := httperr.ParseErrorResponseMessage(res.Body())
		fmt.Println(parsedErr)
		if parsedErr != "" {
			return errors.Wrap(vault.ErrFailedToRememberCipher, parsedErr)
		}
		return errors.Wrap(vault.ErrFailedToRememberCipher, res.String())
	}
	return nil
}
func (r *Provider) RememberCipherCustomBinary(ctx context.Context, c *vault.RememberCipherCustomBinaryData) error {
	r.Lock()
	defer r.Unlock()

	fmt.Println(r.config.Token)
	if r.config.Token.IsExpired() {
		return vault.ErrTokenExpired
	}
	url := r.runAddr + "/custom-binary-data/remember"

	res, err := r.client.
		R().
		SetHeader("Authorization", "Bearer "+r.config.Token.Token).
		SetBody(RememberCipherCustomBinaryDataRequest{
			Key:   c.GetKey(),
			Meta:  c.GetMeta(),
			Value: c.GetValue(),
		}).
		Post(url)
	if err != nil {
		return errors.Wrap(err, "provider: could not create request")
	}

	if res.StatusCode() != http.StatusCreated {
		parsedErr := httperr.ParseErrorResponseMessage(res.Body())
		fmt.Println(parsedErr)
		if parsedErr != "" {
			return errors.Wrap(vault.ErrFailedToRememberCipher, parsedErr)
		}
		return errors.Wrap(vault.ErrFailedToRememberCipher, res.String())
	}
	return nil
}
func (r *Provider) RememberCipherCard(ctx context.Context, c *vault.RememberCipherCardData) error {
	r.Lock()
	defer r.Unlock()

	fmt.Println(r.config.Token)
	if r.config.Token.IsExpired() {
		return vault.ErrTokenExpired
	}
	url := r.runAddr + "/card-data/remember"

	res, err := r.client.
		R().
		SetHeader("Authorization", "Bearer "+r.config.Token.Token).
		SetBody(RememberCipherCardDataRequest{
			Brand:          c.GetBrand(),
			CardHolderName: c.GetCardHolderName(),
			Code:           c.GetCode(),
			ExpMonth:       c.GetExpMonth(),
			ExpYear:        c.GetExpYear(),
			Meta:           c.GetMeta(),
			Number:         c.GetNumber(),
		}).
		Post(url)
	if err != nil {
		return errors.Wrap(err, "provider: could not create request")
	}

	if res.StatusCode() != http.StatusCreated {
		parsedErr := httperr.ParseErrorResponseMessage(res.Body())
		fmt.Println(parsedErr)
		if parsedErr != "" {
			return errors.Wrap(vault.ErrFailedToRememberCipher, parsedErr)
		}
		return errors.Wrap(vault.ErrFailedToRememberCipher, res.String())
	}
	return nil
}
