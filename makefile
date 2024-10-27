.PHONY: mg-create mg_name mg-down mg-remove-tables default run build test docs clean

mg_name ?= migracao_without_name

mg-create:
	migrate create -ext sql -dir internal/infra/db/migrations -seq $(mg_name)

mg-up:
	migrate -path internal/infra/db/migrations -database "mysql://root:root@tcp(localhost:9314)/go_admin_db" up

mg-down:
	migrate -path internal/infra/db/migrations -database "mysql://root:root@tcp(localhost:9314)/go_admin_db" down 1

mg-remove-tables:
	migrate -path internal/infra/db/migrations -database "mysql://root:root@tcp(localhost:9314)/go_admin_db" drop

# API
APP_NAME=API-BACKOFFICE_ZUMENUs

default: run-with-docs

run:
	@go run ./cmd/server/main.go
run-with-docs:
	@swag init -g ./cmd/server/main.go -o ./api
	@go run ./cmd/server/main.go
build:
	@go build -o $(APP_NAME) ./cmd/server/main.go
test:
	@go test ./...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./api