package api

import (
	"github.com/Archetarcher/gophkeeper/internal/client/app"
	"github.com/Archetarcher/gophkeeper/internal/client/app/command"
	"github.com/Archetarcher/gophkeeper/internal/client/app/query"
	"github.com/Archetarcher/gophkeeper/internal/common/auth"
	"github.com/Archetarcher/gophkeeper/internal/common/server/httperr"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"net/http"
)

type HTTPServer struct {
	app app.Application
}

func NewHTTPServer(app app.Application) HTTPServer {
	return HTTPServer{app: app}
}
func (s HTTPServer) RememberCipherCard(w http.ResponseWriter, r *http.Request) {
	rememberCipher := RememberCipherCard{}
	if err := render.Decode(r, &rememberCipher); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}
	cmd := command.RememberCipherCardData{
		CardHolderName: rememberCipher.CardHolderName,
		Brand:          rememberCipher.Brand,
		Number:         rememberCipher.Number,
		ExpYear:        rememberCipher.ExpirationYear,
		ExpMonth:       rememberCipher.ExpirationMonth,
		Code:           rememberCipher.Code,
	}
	cErr := s.app.Commands.RememberCipherCardData.Handle(r.Context(), cmd)
	if cErr != nil {
		httperr.RespondWithSlugError(cErr, w, r)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
func (s HTTPServer) RememberCipherCustomBinary(w http.ResponseWriter, r *http.Request) {
	rememberCipher := RememberCipherCustomBinary{}
	if err := render.Decode(r, &rememberCipher); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}
	cmd := command.RememberCipherCustomBinaryData{
		Key:   rememberCipher.Key,
		Value: rememberCipher.Value,
	}
	cErr := s.app.Commands.RememberCipherCustomBinaryData.Handle(r.Context(), cmd)
	if cErr != nil {
		httperr.RespondWithSlugError(cErr, w, r)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
func (s HTTPServer) RememberCipherCustom(w http.ResponseWriter, r *http.Request) {
	rememberCipher := RememberCipherCustom{}
	if err := render.Decode(r, &rememberCipher); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}

	cmd := command.RememberCipherCustomData{
		Key:   rememberCipher.Key,
		Value: rememberCipher.Value,
	}
	cErr := s.app.Commands.RememberCipherCustomData.Handle(r.Context(), cmd)
	if cErr != nil {
		httperr.RespondWithSlugError(cErr, w, r)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
func (s HTTPServer) RememberCipherLogin(w http.ResponseWriter, r *http.Request) {
	rememberCipher := RememberCipherLogin{}
	if err := render.Decode(r, &rememberCipher); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}

	cmd := command.RememberCipherLoginData{
		Login:    rememberCipher.Login,
		Password: rememberCipher.Password,
		Uri:      rememberCipher.Uri,
	}
	cErr := s.app.Commands.RememberCipherLoginData.Handle(r.Context(), cmd)
	if cErr != nil {
		httperr.RespondWithSlugError(cErr, w, r)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

func (s HTTPServer) SignUp(w http.ResponseWriter, r *http.Request) {
	signUp := SignUp{}
	if err := render.Decode(r, &signUp); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}
	err := s.app.Commands.SignUp.Handle(r.Context(), command.SignUp{
		Login:     signUp.Login,
		Password:  signUp.Password,
		Firstname: signUp.Firstname,
		Lastname:  signUp.Lastname,
	})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
func (s HTTPServer) SignIn(w http.ResponseWriter, r *http.Request) {
	signIn := SignIn{}
	if err := render.Decode(r, &signIn); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}
	token, err := s.app.Queries.SignIn.Handle(r.Context(), query.SignIn{
		Login:    signIn.Login,
		Password: signIn.Password,
	})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	render.Respond(w, r, Token{
		ExpiresAt: token.ExpiresAt,
		Token:     token.Token,
	})
}

// HandlerFromMuxWithJWT creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMuxWithJWT(si ServerInterface, r chi.Router, tokenCfg auth.JWTTokenConfig) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
		Middlewares: []MiddlewareFunc{
			jwtauth.Authenticator(tokenCfg.GetAuthToken()),
			jwtauth.Verifier(tokenCfg.GetAuthToken()),
		},
	})
}

// SignIn defines model for SignIn.
type SignIn struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// SignUp defines model for SignUp.
type SignUp struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Login     string `json:"login"`
	Password  string `json:"password"`
}

// Token defines model for Token.
type Token struct {
	ExpiresAt string `json:"expires_at"`
	Token     string `json:"token"`
}
