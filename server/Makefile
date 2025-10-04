# Load environment variables from .env
include .env
export

# Path to Go-installed migrate binary
MIGRATE := $(HOME)/pkg/bin/migrate

SWAG := $(shell go env GOPATH)/bin/swag
SWAG_SRC := cmd/app/main.go
SWAG_OUT := docs

# Run the application
run:
	go run cmd/app/main.go

# Run migrations
migrate-up:
	go run github.com/pressly/goose/v3/cmd/goose@latest -dir ./migrations postgres "$$DB_URL" up

migrate-down:
	go run github.com/pressly/goose/v3/cmd/goose@latest -dir ./migrations postgres "$$DB_URL" down

swagger:
	@echo "📘 Generating Swagger docs..."
	@$(SWAG) init -g $(SWAG_SRC) -o $(SWAG_OUT)
	@echo "✅ Swagger docs generated in $(SWAG_OUT)/"

swagger-clean:
	@echo "🧹 Removing generated Swagger docs..."
	rm -rf $(SWAG_OUT)