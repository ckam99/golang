test:
  echo "hello"
migrate:
   migrate -source file://internal/database/migrations \
    -database postgres://postgres:postgres@host.docker.internal/golang_echo?sslmode=disable up

migrate-rollback:
   migrate -source file://internal/database/migrations \
    -database postgres://postgres:postgres@host.docker.internal/golang_echo?sslmode=disable down