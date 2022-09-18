MIGRATION_DIR=database/migrations
DATABASE_URL=postgres://postgres:postgres@host.docker.internal:54323/demo?sslmode=disable


up:
	make updb && make createdb && make test

down:
	make dropdb && make downdb

updb:
	docker run --name postgres-dev -p 54323:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres  -d postgres:12-alpine

downdb:
	docker kill postgres-dev

createdb:
	docker exec -it postgres-dev createdb --username=postgres --owner=postgres demo

dropdb:
	docker exec -it postgres-dev dropdb -i -e demo --username=postgres

migrate:
	migrate -source file://${MIGRATION_DIR} \
		-database ${DATABASE_URL} up

migratedown:
	migrate -source file://${MIGRATION_DIR} \
		-database ${DATABASE_URL} down

createmigration:
	migrate create -ext sql -dir database/migrations -seq $(name)

sqlc:
	sqlc generate


test:
	go test  -v -cover ./...

.PHONY: migrate migratedown createmigration sqlc test updb downdb createdb dropdb