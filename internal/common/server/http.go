package server

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/Archetarcher/gophkeeper/internal/common/auth"
	"github.com/Archetarcher/gophkeeper/internal/common/encryption"
	"github.com/Archetarcher/gophkeeper/internal/common/server/httperr"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func RunHTTPServerOnAddr(addr string, handler func(router chi.Router) http.Handler) {
	apiRouter := chi.NewRouter()
	rootRouter := chi.NewRouter()

	// we are mounting all APIs under /api path
	rootRouter.Mount("/api", handler(apiRouter))

	logrus.Info("Starting HTTP server", addr)
	err := http.ListenAndServe(addr, rootRouter)
	if err != nil {
		logrus.WithError(err).Panic("Unable to start HTTP server")
	}

}

func RunHTTPServerOnAddrWithMiddlewares(addr string, handler func(router chi.Router) http.Handler, serverConfig *Config, tokenCfg auth.JWTTokenConfig) {
	// config middlewares only for api routes
	// we are mounting all APIs under /api path
	rootRouter := chi.NewRouter()
	rootRouter.Use(jwtauth.Verifier(tokenCfg.GetAuthToken()))
	rootRouter.Use(jwtauth.Authenticator(tokenCfg.GetAuthToken()))

	rootRouter.Post("/session", serverConfig.handleStartSession)

	rootRouter.Mount("/api", apiRouter(handler, serverConfig))
	logrus.Info("Starting HTTP server", addr)
	err := http.ListenAndServe(addr, rootRouter)
	if err != nil {
		logrus.WithError(err).Panic("Unable to start HTTP server")
	}

}
func apiRouter(handler func(router chi.Router) http.Handler, serverConfig *Config) http.Handler {
	router := chi.NewRouter()
	//router.Use(GzipMiddleware)
	//router.Use(func(handler http.Handler) http.Handler {
	//	enc := encryption.NewSymmetric(serverConfig.Session.Key)
	//	return RequestDecryptMiddleware(handler, enc)
	//})
	return handler(router)
}

func (c *Config) handleStartSession(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("start session request")
	startSession := StartSession{}
	if err := render.Decode(request, &startSession); err != nil {
		httperr.BadRequest("invalid-request", err, writer, request)
		return
	}
	_, err := auth.GetIDFromToken(request.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, writer, request)
		return
	}
	sDec, err := b64.StdEncoding.DecodeString(startSession.Key)
	if err != nil {
		httperr.RespondWithSlugError(err, writer, request)
		return
	}
	key, err := encryption.NewAsymmetric(os.Getenv("PUBLIC_KEY_PATH"), os.Getenv("PRIVATE_KEY_PATH")).Decrypt(sDec)
	if err != nil {
		httperr.RespondWithSlugError(err, writer, request)
		return
	}
	c.Session.Key = string(key)
	writer.WriteHeader(http.StatusOK)
}

// StartSession defines model for StartSession.
type StartSession struct {
	Key string `json:"key"`
}
