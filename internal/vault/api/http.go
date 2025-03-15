package api

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Archetarcher/gophkeeper/internal/common/auth"
	"github.com/Archetarcher/gophkeeper/internal/common/encryption"
	"github.com/Archetarcher/gophkeeper/internal/common/server"
	"github.com/Archetarcher/gophkeeper/internal/common/server/httperr"
	"github.com/Archetarcher/gophkeeper/internal/vault/app"
	"github.com/Archetarcher/gophkeeper/internal/vault/app/command"
	"github.com/Archetarcher/gophkeeper/internal/vault/app/query"
	"github.com/go-chi/render"
	"net/http"
	"os"
)

type HTTPServer struct {
	app    app.Application
	config *server.Config
}

func NewHTTPServer(app app.Application, config *server.Config) HTTPServer {
	return HTTPServer{app: app, config: config}
}
func (s HTTPServer) StartSession(w http.ResponseWriter, r *http.Request) {
	startSession := StartSession{}
	if err := json.NewDecoder(r.Body).Decode(&startSession); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}
	sDec, err := b64.StdEncoding.DecodeString(startSession.Key)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	key, err := encryption.NewAsymmetric(os.Getenv("PUBLIC_KEY_PATH"), os.Getenv("PRIVATE_KEY_PATH")).Decrypt(sDec)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	fmt.Println(string(key))
	s.config.Session.Key = string(key)
	w.WriteHeader(http.StatusOK)
}

// StartSession defines model for StartSession.
type StartSession struct {
	Key string `json:"key"`
}

func (s HTTPServer) RememberCipherLoginData(w http.ResponseWriter, r *http.Request) {
	rememberCipher := RememberCipherLoginData{}
	if err := json.NewDecoder(r.Body).Decode(&rememberCipher); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}
	id, err := auth.GetIDFromToken(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	cmd := command.RememberCipherLoginData{
		Login:    rememberCipher.Login,
		Password: rememberCipher.Password,
		Uri:      rememberCipher.Uri,
		UserId:   id,
		Meta:     rememberCipher.Meta,
	}
	cErr := s.app.Commands.RememberCipherLoginData.Handle(r.Context(), cmd)
	if cErr != nil {
		httperr.RespondWithSlugError(cErr, w, r)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
func (s HTTPServer) RememberCipherCustomData(w http.ResponseWriter, r *http.Request) {
	rememberCipher := RememberCipherCustomData{}
	if err := json.NewDecoder(r.Body).Decode(&rememberCipher); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}
	id, err := auth.GetIDFromToken(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	cmd := command.RememberCipherCustomData{
		Key:    rememberCipher.Key,
		Value:  rememberCipher.Value,
		Meta:   rememberCipher.Meta,
		UserId: id,
	}
	cErr := s.app.Commands.RememberCipherCustomData.Handle(r.Context(), cmd)
	if cErr != nil {
		httperr.RespondWithSlugError(cErr, w, r)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
func (s HTTPServer) RememberCipherCustomBinaryData(w http.ResponseWriter, r *http.Request) {
	rememberCipher := RememberCipherCustomBinaryData{}
	if err := json.NewDecoder(r.Body).Decode(&rememberCipher); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}
	id, err := auth.GetIDFromToken(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	cmd := command.RememberCipherCustomBinaryData{
		Key:    rememberCipher.Key,
		Value:  rememberCipher.Value,
		Meta:   rememberCipher.Meta,
		UserId: id,
	}
	cErr := s.app.Commands.RememberCipherCustomBinaryData.Handle(r.Context(), cmd)
	if cErr != nil {
		httperr.RespondWithSlugError(cErr, w, r)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
func (s HTTPServer) RememberCipherCardData(w http.ResponseWriter, r *http.Request) {
	rememberCipher := RememberCipherCardData{}
	if err := json.NewDecoder(r.Body).Decode(&rememberCipher); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}
	id, err := auth.GetIDFromToken(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	cmd := command.RememberCipherCardData{
		CardHolderName: rememberCipher.CardHolderName,
		Brand:          rememberCipher.Brand,
		Number:         rememberCipher.Number,
		ExpYear:        rememberCipher.ExpYear,
		ExpMonth:       rememberCipher.ExpMonth,
		Code:           rememberCipher.Code,
		Meta:           rememberCipher.Meta,
		UserId:         id,
	}
	cErr := s.app.Commands.RememberCipherCardData.Handle(r.Context(), cmd)
	if cErr != nil {
		httperr.RespondWithSlugError(cErr, w, r)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

func (s HTTPServer) ShowUserSecrets(w http.ResponseWriter, r *http.Request) {
	id, err := auth.GetIDFromToken(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	userSecrets, err := s.app.Queries.ShowUserSecrets.Handle(r.Context(), query.ShowUserSecrets{
		UserId: id,
	})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	secrets := userSecretsToResponse(userSecrets)
	secretsResponse := Secrets{secrets}

	render.Respond(w, r, secretsResponse)
}

func (s HTTPServer) ShowSecret(w http.ResponseWriter, r *http.Request) {
	showSecret := ShowSecret{}
	if err := render.Decode(r, &showSecret); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}
	id, err := auth.GetIDFromToken(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	userSecret, err := s.app.Queries.ShowSecret.Handle(r.Context(), query.ShowSecret{
		UserId: id,
		Key:    showSecret.Key,
	})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}
	secretResponse := Secret{
		Data: userSecret.Data,
		Key:  userSecret.Key,
		Type: userSecret.SecretType,
		Uuid: userSecret.ID,
	}

	render.Respond(w, r, secretResponse)
}
func userSecretsToResponse(userSecrets []query.Secret) []Secret {
	var secrets []Secret
	for _, tm := range userSecrets {
		t := Secret{
			Uuid: tm.ID,
			Key:  tm.Key,
			Data: tm.Data,
			Type: tm.SecretType,
		}

		secrets = append(secrets, t)
	}

	return secrets
}
