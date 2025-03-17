package api

import (
	"github.com/Archetarcher/gophkeeper/internal/auth/app"
	"github.com/Archetarcher/gophkeeper/internal/auth/app/command"
	"github.com/Archetarcher/gophkeeper/internal/auth/app/query"
	"github.com/Archetarcher/gophkeeper/internal/common/server/httperr"
	"github.com/go-chi/render"
	"net/http"
)

type HTTPServer struct {
	app app.Application
}

func NewHTTPServer(app app.Application) HTTPServer {
	return HTTPServer{app: app}
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

func (s HTTPServer) GetCurrentUser(w http.ResponseWriter, r *http.Request) {

}
