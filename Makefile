VAULT_BINARY_NAME=vault
CLIENT_BINARY_NAME=client
AUTH_BINARY_NAME=auth


client-build:
	cd cmd/client &&  go build -ldflags "-X main.buildVersion=1.0.0 -X main.buildDate=$(date +%Y-%m-%d) -X main.buildCommit=$(git rev-parse HEAD)" -o ${CLIENT_BINARY_NAME}

client-build-platforms:
	cd cmd/client && GOARCH=amd64 GOOS=darwin go build -o ${CLIENT_BINARY_NAME}-darwin main.go
	cd cmd/client && GOARCH=amd64 GOOS=linux go build -o ${CLIENT_BINARY_NAME}-linux main.go
	cd cmd/client && GOARCH=amd64 GOOS=windows go build -o ${CLIENT_BINARY_NAME}-windows main.go

client-run: client-build
	./cmd/client/${CLIENT_BINARY_NAME}



auth-build:
	cd cmd/auth &&  go build -ldflags "-X main.buildVersion=1.0.0 -X main.buildDate=$(date +%Y-%m-%d) -X main.buildCommit=$(git rev-parse HEAD)" -o ${AUTH_BINARY_NAME}

auth-build-platforms:
	cd cmd/auth && GOARCH=amd64 GOOS=darwin go build -o ${AUTH_BINARY_NAME}-darwin main.go
	cd cmd/auth && GOARCH=amd64 GOOS=linux go build -o ${AUTH_BINARY_NAME}-linux main.go
	cd cmd/auth && GOARCH=amd64 GOOS=windows go build -o ${AUTH_BINARY_NAME}-windows main.go

auth-run: auth-build
	./cmd/auth/${AUTH_BINARY_NAME}


vault-build:
	cd cmd/vault &&  go build -ldflags "-X main.buildVersion=1.0.0 -X main.buildDate=$(date +%Y-%m-%d) -X main.buildCommit=$(git rev-parse HEAD)" -o ${VAULT_BINARY_NAME}

vault-build-platforms:
	cd cmd/vault && GOARCH=amd64 GOOS=darwin go build -o ${VAULT_BINARY_NAME}-darwin main.go
	cd cmd/vault && GOARCH=amd64 GOOS=linux go build -o ${VAULT_BINARY_NAME}-linux main.go
	cd cmd/vault && GOARCH=amd64 GOOS=windows go build -o ${VAULT_BINARY_NAME}-windows main.go

vault-run: vault-build
	./cmd/vault/${VAULT_BINARY_NAME} -d=postgres://postgres:postgres@localhost:5432/praktikum?sslmode=disable


test:
	go test ./...

test_coverage:
	go test ./... -coverprofile cover.out && go tool cover -func cover.out


vet:
	go vet

