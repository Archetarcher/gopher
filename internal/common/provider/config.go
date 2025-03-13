package provider

import (
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"time"
)

type Config struct {
	RunAddr string
	Client  *resty.Client
	Token   *Token
	Session *Session
}

// Token is a struct for obtained jwt token
type Token struct {
	Token     string
	ExpiresAt string
}

func (t *Token) IsExpired() bool {
	expires, err := time.Parse(time.RFC3339, t.ExpiresAt)
	if err != nil {
		logrus.Errorf("failed to parse token expiration time %p", err)
		return true
	}
	return expires.Unix() < time.Now().Unix()
}

// Session is a struct for obtained session
type Session struct {
	Key string
}
