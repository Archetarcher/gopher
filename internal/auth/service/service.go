package service

import (
	"context"
	"github.com/Archetarcher/gophkeeper/internal/auth/app"
	"github.com/Archetarcher/gophkeeper/internal/auth/app/command"
	"github.com/Archetarcher/gophkeeper/internal/auth/app/query"
	"github.com/Archetarcher/gophkeeper/internal/auth/domain/user/pgx"
	"github.com/Archetarcher/gophkeeper/internal/common/auth"
	"github.com/Archetarcher/gophkeeper/internal/common/db"
	"github.com/pkg/errors"
	"log"
	"os"
)

func NewApplication(ctx context.Context, tokenCfg auth.JWTTokenConfig) app.Application {
	userRepository, err := pgx.New(ctx, db.Config{
		Dsn:            os.Getenv("PGX_DSN"),
		MigrationsPath: os.Getenv("PGX_AUTH_MIGRATIONS_PATH"),
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to init db"))
	}
	return app.Application{
		Commands: app.Commands{
			SignUp: command.NewSignUpHandler(userRepository),
		},
		Queries: app.Queries{
			SignIn: query.NewSignInHandler(userRepository, tokenCfg),
		},
	}
}
