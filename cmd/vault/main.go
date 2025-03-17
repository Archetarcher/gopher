package main

import (
	"context"
	"github.com/Archetarcher/gophkeeper/internal/common/auth"
	"github.com/Archetarcher/gophkeeper/internal/common/server"
	"github.com/Archetarcher/gophkeeper/internal/vault/api"
	"github.com/Archetarcher/gophkeeper/internal/vault/service"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	// jwt config
	jwtTokenCfg := auth.GetNewJWTTokenConfig()
	// new application
	app := service.NewApplication(ctx)

	// server side config
	serverConfig := &server.Config{Session: &server.Session{}}

	// new application server
	s := api.NewHTTPServer(app, serverConfig)

	// some root routes that exceeds basic middlewares for app
	rootRouter := chi.NewRouter()
	rootRouter.Post("/session", s.StartSession)

	// run server
	server.RunHTTPServerOnAddrWithMiddlewares(":"+os.Getenv("VAULT_PORT"), func(router chi.Router) http.Handler {
		return api.HandlerFromMux(s, router)
	}, rootRouter,
		jwtauth.Verifier(jwtTokenCfg.GetAuthToken()),
		jwtauth.Authenticator(jwtTokenCfg.GetAuthToken()),
		func(handler http.Handler) http.Handler {
			return server.RequestDecryptMiddleware(handler, serverConfig)
		},
		server.GzipMiddleware,
		middleware.DefaultLogger,
	)
}
