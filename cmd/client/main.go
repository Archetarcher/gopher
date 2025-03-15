package main

import (
	"github.com/Archetarcher/gophkeeper/internal/client/api"
	"github.com/Archetarcher/gophkeeper/internal/client/service"
	"github.com/Archetarcher/gophkeeper/internal/common/auth"
	"github.com/Archetarcher/gophkeeper/internal/common/provider"
	"github.com/Archetarcher/gophkeeper/internal/common/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	jwtTokenCfg := auth.GetNewJWTTokenConfig()
	prvConfig := &provider.Config{
		Token:   &provider.Token{},
		Session: &provider.Session{},
	}
	app := service.NewApplication(ctx, prvConfig)
	s := api.NewHTTPServer(app)
	rootRouter := chi.NewRouter()
	rootRouter.Post("/sign-up", s.SignUp)
	rootRouter.Post("/sign-in", s.SignIn)
	server.RunHTTPServerOnAddrWithMiddlewares(":"+os.Getenv("CLIENT_PORT"), func(router chi.Router) http.Handler {
		return api.HandlerFromMux(s, router)
	}, rootRouter,
		jwtauth.Verifier(jwtTokenCfg.GetAuthToken()),
		jwtauth.Authenticator(jwtTokenCfg.GetAuthToken()),
		middleware.DefaultLogger,
		func(handler http.Handler) http.Handler {
			return provider.CheckTokenAuthority(handler, prvConfig, jwtTokenCfg.GetAuthToken())
		},
	)
}
