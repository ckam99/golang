BINARY_NAME=server.out
DB_DRIVER=postgres
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=golang_echo
DB_HOST=host.docker.internal
DB_PORT=5432
DB_SSLMODE=disable
TIMEZONE=Europe/Moscow

DATABASE_URL=${DB_DRIVER}://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}
MIGRATION_DIR=internal/database/postgres/migrations

dev:
	go run cmd/server.go

build:
	go build -o ${BINARY_NAME} cmd/server.go

run:
	go build -o ${BINARY_NAME} cmd/server.go
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}

migrate:
	migrate -source file://${MIGRATION_DIR}  -database ${DATABASE_URL} up

migratedown:
	migrate -source file://${MIGRATION_DIR}  -database ${DATABASE_URL} down
