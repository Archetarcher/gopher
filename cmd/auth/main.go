package main

import (
	"context"
	"github.com/Archetarcher/gophkeeper/internal/auth/api"
	"github.com/Archetarcher/gophkeeper/internal/auth/service"
	"github.com/Archetarcher/gophkeeper/internal/common/auth"
	"github.com/Archetarcher/gophkeeper/internal/common/server"
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
	app := service.NewApplication(ctx, jwtTokenCfg)

	// run application server
	server.RunHTTPServerOnAddr(":"+os.Getenv("AUTH_PORT"), func(router chi.Router) http.Handler {
		return api.HandlerFromMux(api.NewHTTPServer(app), router)
	})
}
