package auth

import (
	"context"
	"github.com/go-chi/jwtauth/v5"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"os"
	"strconv"
	"strings"
	"time"
)

type JWTTokenConfig struct {
	key              string
	expiresInMinutes int
	authToken        *jwtauth.JWTAuth
}

func GetNewJWTTokenConfig() JWTTokenConfig {
	key := strings.ToLower(os.Getenv("SERVER_TOKEN_KEY"))
	expiration := os.Getenv("SERVER_TOKEN_EXPIRATION_IN_MINUTES")

	return setUpJWTTokenConfig(key, expiration)
}

func (t *JWTTokenConfig) CreateToken(id uuid.UUID) (string, error) {
	claims := jwt2.MapClaims{
		"id": id,
	}

	jwtauth.SetExpiry(claims, time.Now().Add(time.Minute*time.Duration(t.expiresInMinutes)))
	_, token, err := t.authToken.Encode(claims)
	if err != nil {
		return "", err
	}

	return token, nil
}
func (t *JWTTokenConfig) GetTokenExpiration() string {
	return time.Now().Add(time.Minute * time.Duration(t.expiresInMinutes)).Format(time.RFC3339)
}
func (t *JWTTokenConfig) GetAuthToken() *jwtauth.JWTAuth {
	return t.authToken
}
func IsAuthorized(requestContext context.Context) bool {
	token, _, err := jwtauth.FromContext(requestContext)

	if err != nil {
		return false
	}

	if token != nil && jwt.Validate(token) == nil {
		return true
	}

	return false
}

func GetIDFromToken(requestContext context.Context) (uuid.UUID, error) {
	_, claims, err := jwtauth.FromContext(requestContext)
	if err != nil {
		return uuid.Nil, err
	}
	userIDFromClaims := uuid.MustParse(claims["id"].(string))

	return userIDFromClaims, nil

}
func generateAuthToken(key string) *jwtauth.JWTAuth {
	return jwtauth.New("HS256", []byte(key), nil)
}

func setUpJWTTokenConfig(key, expiration string) JWTTokenConfig {
	tokenExpiration := 0
	if i, err := strconv.Atoi(expiration); err == nil {
		tokenExpiration = i
	}
	return JWTTokenConfig{
		key:              key,
		expiresInMinutes: tokenExpiration,
		authToken:        generateAuthToken(key),
	}

}
