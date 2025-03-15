package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Archetarcher/gophkeeper/internal/client/provider/auth"
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

type SignInRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type SignUpRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Login     string `json:"login"`
	Password  string `json:"password"`
}

func (r *Provider) SignUp(ctx context.Context, u *auth.SignUp) error {
	r.Lock()
	defer r.Unlock()

	url := r.runAddr + "/users/sign-up"

	res, err := r.client.
		R().
		SetBody(SignUpRequest{
			Firstname: u.GetFirstname(),
			Lastname:  u.GetLastname(),
			Login:     u.GetLogin(),
			Password:  u.GetPassword(),
		}).
		Post(url)
	if err != nil {
		return errors.Wrap(err, "provider: could not create request")
	}

	if res.StatusCode() != http.StatusCreated {
		parsedErr := httperr.ParseErrorResponseMessage(res.Body())
		fmt.Println(parsedErr)
		if parsedErr != "" {
			return errors.Wrap(auth.ErrSignUp, parsedErr)
		}
		return errors.Wrap(auth.ErrSignUp, res.String())
	}
	return nil
}

func (r *Provider) SignIn(ctx context.Context, u *auth.SignIn) (*provider.Token, error) {
	r.Lock()
	defer r.Unlock()
	url := r.runAddr + "/users/sign-in"
	res, err := r.client.
		R().
		SetBody(SignInRequest{
			Login:    u.GetLogin(),
			Password: u.GetPassword(),
		}).
		Post(url)
	if err != nil {
		return nil, errors.Wrap(err, "provider: could not create request")
	}

	if res.StatusCode() != http.StatusOK {
		parsedErr := httperr.ParseErrorResponseMessage(res.Body())
		if parsedErr != "" {
			return nil, errors.Wrap(auth.ErrSignIn, parsedErr)
		}
		return nil, errors.Wrap(auth.ErrSignIn, res.String())
	}
	fmt.Println(res.String())
	var token provider.Token
	err = json.Unmarshal(res.Body(), &token)
	if err != nil {
		return nil, errors.Wrap(err, "provider: failed to serialize response into token struct")
	}
	// set bearer token for provider requests
	r.config.Token = &token
	fmt.Println("r.config.Token in auth")
	fmt.Println(r.config.Token)
	return &token, nil
}
