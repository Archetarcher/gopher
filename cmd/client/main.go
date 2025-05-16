package main

import (
	"fmt"
	"github.com/Archetarcher/gophkeeper/internal/client/api"
	"github.com/Archetarcher/gophkeeper/internal/client/cli"
	"github.com/Archetarcher/gophkeeper/internal/client/cli/gophkeeper/commands"
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
	"strings"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

func main() {
	printBuildData()
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	// config for jwt token
	jwtTokenCfg := auth.GetNewJWTTokenConfig()
	// config for providers
	prvConfig := &provider.Config{
		Token:   &provider.Token{},
		Session: &provider.Session{},
	}

	// new application server
	app := service.NewApplication(ctx, prvConfig)
	serverType := strings.ToLower(os.Getenv("SERVER_TO_RUN"))
	switch serverType {
	case "http":
		s := api.NewHTTPServer(app)

		// some root routes that exceeds basic middlewares for app
		rootRouter := chi.NewRouter()
		rootRouter.Post("/sign-up", s.SignUp)
		rootRouter.Post("/sign-in", s.SignIn)
		// run server
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
	case "cli":
		s := commands.NewCLI(app)
		cli.Execute(ctx, s)
	default:
		panic(fmt.Sprintf("server type '%s' is not supported", serverType))
	}

}

func printBuildData() {
	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build commit: %s\n", buildCommit)
}
