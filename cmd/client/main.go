package main

import (
	"github.com/Archetarcher/gophkeeper/internal/client/api"
	"github.com/Archetarcher/gophkeeper/internal/client/service"
	"github.com/Archetarcher/gophkeeper/internal/common/auth"
	"github.com/Archetarcher/gophkeeper/internal/common/server"
	"github.com/go-chi/chi/v5"
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
	app := service.NewApplication(ctx)
	server.RunHTTPServerOnAddr(":"+os.Getenv("CLIENT_PORT"), func(router chi.Router) http.Handler {
		return api.HandlerFromMuxWithJWT(api.NewHTTPServer(app), router, jwtTokenCfg)
	})
}
