package provider

import (
	"github.com/Archetarcher/gophkeeper/internal/common/auth"
	"github.com/Archetarcher/gophkeeper/internal/common/server/httperr"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"net/http"
)

func CheckTokenAuthority(next http.Handler, prvConfig *Config, jwtToken *jwtauth.JWTAuth) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		requestUserId, err := auth.GetIDFromToken(r.Context())
		if err != nil {
			httperr.RespondWithSlugError(err, rw, r)
			return
		}
		token, err := jwtToken.Decode(prvConfig.Token.Token)
		if err != nil {
			httperr.RespondWithSlugError(err, rw, r)
			return
		}
		configUserId, ok := token.Get("id")
		if !ok {
			httperr.RespondWithSlugError(errors.New("token not found in vault config"), rw, r)
			return
		}
		if uuid.MustParse(configUserId.(string)) != requestUserId {
			httperr.RespondWithSlugError(errors.New("authorized user's token and provided user token not matching"), rw, r)
			return
		}

		next.ServeHTTP(rw, r.WithContext(r.Context()))
	})
}
