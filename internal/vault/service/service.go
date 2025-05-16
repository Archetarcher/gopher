package service

import (
	"context"
	"github.com/Archetarcher/gophkeeper/internal/common/db"
	"github.com/Archetarcher/gophkeeper/internal/common/encryption"
	"github.com/Archetarcher/gophkeeper/internal/vault/app"
	"github.com/Archetarcher/gophkeeper/internal/vault/app/command"
	cipherCard "github.com/Archetarcher/gophkeeper/internal/vault/domain/cipherCardData/pgx"
	cipherCustomBinary "github.com/Archetarcher/gophkeeper/internal/vault/domain/cipherCustomBinaryData/pgx"
	cipherCustom "github.com/Archetarcher/gophkeeper/internal/vault/domain/cipherCustomData/pgx"
	cipherLogin "github.com/Archetarcher/gophkeeper/internal/vault/domain/cipherLoginData/pgx"
	"github.com/pkg/errors"
	"log"
	"os"
)

func NewApplication(ctx context.Context) app.Application {
	cfg := db.Config{Dsn: os.Getenv("PGX_DSN"), MigrationsPath: os.Getenv("PGX_VAULT_MIGRATIONS_PATH")}
	cipherLoginRepository, err := cipherLogin.New(ctx, cfg)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to init cipherLoginRepository"))
	}
	cipherCustomRepository, err := cipherCustom.New(ctx, cfg)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to init cipherCustomRepository"))
	}
	cipherCustomBinaryRepository, err := cipherCustomBinary.New(ctx, cfg)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to init cipherCustomBinaryRepository"))
	}
	cipherCardRepository, err := cipherCard.New(ctx, cfg)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to init cipherCustomBinaryRepository"))
	}
	asymmetricEncryption := encryption.NewAsymmetric(os.Getenv("PUBLIC_KEY_PATH"), os.Getenv("PRIVATE_KEY_PATH"))
	return app.Application{
		Commands: app.Commands{
			RememberCipherLoginData:        command.NewRememberCipherLoginDataHandler(cipherLoginRepository, asymmetricEncryption),
			RememberCipherCustomData:       command.NewRememberCipherCustomDataHandler(cipherCustomRepository, asymmetricEncryption),
			RememberCipherCustomBinaryData: command.NewRememberCipherCustomBinaryDataHandler(cipherCustomBinaryRepository, asymmetricEncryption),
			RememberCipherCardData:         command.NewRememberCipherCardDataHandler(cipherCardRepository, asymmetricEncryption),
		},
		Queries: app.Queries{},
	}
}
