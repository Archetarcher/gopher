package service

import (
	"context"
	"github.com/Archetarcher/gophkeeper/internal/client/app"
	"github.com/Archetarcher/gophkeeper/internal/client/app/command"
	"github.com/Archetarcher/gophkeeper/internal/client/app/query"
	auth "github.com/Archetarcher/gophkeeper/internal/client/provider/auth/http"
	vault "github.com/Archetarcher/gophkeeper/internal/client/provider/vault/http"
	"github.com/Archetarcher/gophkeeper/internal/common/encryption"
	"github.com/Archetarcher/gophkeeper/internal/common/provider"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"log"
	"os"
)

func NewApplication(ctx context.Context) app.Application {
	vaultAddr := os.Getenv("VAULT_RUN_ADDR")
	authAddr := os.Getenv("AUTH_RUN_ADDR")

	prvConfig := &provider.Config{
		Client:  resty.New(),
		RunAddr: "http://",
		Token:   &provider.Token{},
		Session: &provider.Session{},
	}
	authProvider := auth.New(prvConfig, authAddr)
	vaultProvider := vault.New(prvConfig, vaultAddr)

	asymmetricEncryption := encryption.NewAsymmetric(os.Getenv("PUBLIC_KEY_PATH"), os.Getenv("PRIVATE_KEY_PATH"))
	err := vaultProvider.StartSession(ctx, asymmetricEncryption)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to start session"))
	}
	symmetricEncryption := encryption.NewSymmetric(prvConfig.Session.Key)

	prvConfig.Client.OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
		return provider.GzipAndEncryptMiddleware(client, request, symmetricEncryption)
	})

	return app.Application{
		Commands: app.Commands{
			RememberCipherLoginData:        command.NewRememberCipherLoginDataHandler(vaultProvider),
			RememberCipherCustomData:       command.NewRememberCipherCustomDataHandler(vaultProvider),
			RememberCipherCustomBinaryData: command.NewRememberCipherCustomBinaryDataHandler(vaultProvider),
			RememberCipherCardData:         command.NewRememberCipherCardDataHandler(vaultProvider),
			SignUp:                         command.NewSignUpHandler(authProvider),
		},
		Queries: app.Queries{
			SignIn: query.NewSignInHandler(authProvider),
		},
	}
}
