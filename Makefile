MIGRATION_DIR=internal/provider/postgres/migrations
DATABASE_URL=postgres://postgres:postgres@host.docker.internal/demo?sslmode=disable
PROTO_OUT=internal/controller/rpc/pb
PROTO_PATH=internal/controller/rpc/proto

pb:
	rm -rf ${PROTO_OUT}/*.go
	protoc --proto_path=${PROTO_PATH} \
	--go_out=${PROTO_OUT} \
	--go_opt=paths=source_relative \
	--go-grpc_out=${PROTO_OUT} \
	--go-grpc_opt=paths=source_relative ${PROTO_PATH}/*.proto

migration:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir ${MIGRATION_DIR} -seq $$name

migrate:
	migrate -source file://${MIGRATION_DIR} -database ${DATABASE_URL} up

rollback:
	migrate -source file://${MIGRATION_DIR} -database ${DATABASE_URL} down

evans:
	evans --host localhost --port 5000 -r repl
serve:
	go run cmd/main.go
