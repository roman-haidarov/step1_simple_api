DB_DSN := "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

.PHONY: gen
gen:
	mkdir -p ./internal/web/tasks
	mkdir -p ./internal/web/users

	oapi-codegen -generate chi-server,types -package tasks -include-tags tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go
	oapi-codegen -generate chi-server,types -package users -include-tags users openapi/openapi.yaml > ./internal/web/users/api.gen.go

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

run:
	go run cmd/main.go

lint:
	golangci-lint run --out-format=colored-line-number
