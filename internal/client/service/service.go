package service

import (
	"context"
	b64 "encoding/base64"
	"fmt"
	"github.com/Archetarcher/gophkeeper/internal/client/app"
	"github.com/Archetarcher/gophkeeper/internal/client/app/command"
	"github.com/Archetarcher/gophkeeper/internal/client/app/query"
	auth "github.com/Archetarcher/gophkeeper/internal/client/provider/auth/http"
	vault "github.com/Archetarcher/gophkeeper/internal/client/provider/vault/http"
	"github.com/Archetarcher/gophkeeper/internal/common/encryption"
	"github.com/Archetarcher/gophkeeper/internal/common/provider"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
)

func NewApplication(ctx context.Context, prvConfig *provider.Config) app.Application {

	// new auth provider
	authProvider := auth.New(prvConfig, os.Getenv("AUTH_RUN_ADDR"))

	asymmetricEncryption := encryption.NewAsymmetric(os.Getenv("PUBLIC_KEY_PATH"), os.Getenv("PRIVATE_KEY_PATH"))
	// start session with vault
	err := startVaultSession(asymmetricEncryption, prvConfig)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to start session"))
	}
	// new vault provider with middlewares
	vaultProvider := vault.New(prvConfig, os.Getenv("VAULT_RUN_ADDR"),
		func(client *resty.Client, request *resty.Request) error {
			return provider.GzipAndEncryptMiddleware(client, request, prvConfig)
		},
	)

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
func startVaultSession(enc encryption.AsymmetricEncryption, prvConfig *provider.Config) error {
	url := os.Getenv("VAULT_RUN_ADDR") + "/session"
	key, gErr := encryption.GenKey(16)
	if gErr != nil {
		return errors.Wrap(gErr, "provider: failed to generate crypto key")
	}
	encryptedKey, eErr := enc.Encrypt(key)
	if eErr != nil {
		return eErr
	}

	sEnc := b64.StdEncoding.EncodeToString(encryptedKey)

	res, err := resty.New().
		R().
		SetBody(map[string]string{
			"key": sEnc,
		}).
		Post(url)

	if err != nil {
		return errors.Wrap(err, "provider: could not create request")
	}

	if res.StatusCode() != http.StatusOK {
		return fmt.Errorf("failed to start session with vault %s", res.String())
	}

	logrus.Info("<<<<<<<<<starting vault session>>>>>>>>>>")
	fmt.Println(string(key))
	prvConfig.Session.Key = string(key)
	return nil
}
