package http

import (
	"context"
	"encoding/json"
	"github.com/Archetarcher/gophkeeper/internal/client/provider/auth"
	"github.com/Archetarcher/gophkeeper/internal/common/provider"
	"github.com/go-resty/resty/v2"
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
	config.Client = resty.New()
	return &Provider{
		config: config,
	}
}

func (r *Provider) SignUp(ctx context.Context, u *auth.SignUp) error {
	r.Lock()
	defer r.Unlock()

	url := r.config.RunAddr + "/users/sign-up"

	res, err := r.config.Client.
		R().
		SetBody(u).
		Post(url)
	if err != nil {
		return errors.Wrap(err, "provider: could not create request")
	}

	if res.StatusCode() != http.StatusCreated {
		return errors.Wrap(err, "provider: responded with error")

	}
	return nil
}

func (r *Provider) SignIn(ctx context.Context, u *auth.SignIn) (*provider.Token, error) {
	r.Lock()
	defer r.Unlock()

	url := r.config.RunAddr + "/users/sign-in"

	res, err := r.config.Client.
		R().
		SetBody(u).
		Post(url)
	if err != nil {
		return nil, errors.Wrap(err, "provider: could not create request")
	}

	if res.StatusCode() != http.StatusOK {
		return nil, errors.Wrap(err, "provider: responded with error")
	}
	var token provider.Token
	err = json.Unmarshal([]byte(res.Body()), &token)
	if err != nil {
		return nil, errors.Wrap(err, "provider: failed to serialize response into token struct")
	}
	// set bearer token for provider requests
	r.config.Token = &token
	return &token, nil
}
