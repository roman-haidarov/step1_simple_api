DB_DSN := "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

.PHONY: gen
gen:
	oapi-codegen -generate chi-server,types -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

run:
	go run cmd/main.go 
