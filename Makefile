MIGRATION_DIR=database/migrations
DATABASE_URL=postgres://postgres:postgres@host.docker.internal/golang_sqlx?sslmode=disable

migrate:
	migrate -source file://${MIGRATION_DIR} \
		-database ${DATABASE_URL} up

migratedown:
	migrate -source file://${MIGRATION_DIR} \
		-database ${DATABASE_URL} down

createmigration:
	migrate create -ext sql -dir database/migrations -seq $(name)

build:
	go build -o serve main.go

start:
	./serve

dev:
	go run main.go