# Makefile for Go project
API="api"
RUN_PATH="./cmd/api/main.go"

# Build the Go project
run:
	@go run $(RUN_PATH)

local-setup:
	docker build -t go-exinity-task-postgress -f ./scripts/database/Dockerfile ./scripts/database/
	docker run --name go-exinity-task-postgress -p 5432:5432 -d go-exinity-task-postgress

docker-compose:
	docker-compose up

generate-grpc:
	@protoc --go_out=./internal/app/proto  --go-grpc_out=./internal/app/proto ./internal/app/proto/candle.proto
