.PHONY: build dev migrate-dev docs mockgen

build:
	go build -o ./app ./cmd/main.go

dev:
	air -c ./.air.toml

migrate-dev:
	go run cmd/database/migrate/main.go -db-socket tcp -db-host localhost -db-port 3306 -db-name clean-architecture -db-username root -db-password root

docs:
	swag init -d ./internal/interfaces/http -g ./handler/handler.go --pd -o ./docs/

mockgen:
	go install go.uber.org/mock/mockgen@latest
	rm -rf ./mock
	go generate ./...