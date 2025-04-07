# Makefile for Go project
API="api"
RUN_PATH="./cmd/api/main.go"
SWAGGER_URL="http://localhost:1989/swagger/index.html"
PORT="1989"

# Build the Go project
build: clear
	@go build -o $(API) $(RUN_PATH)

clear:
	@rm -f ./$(API)

run:
	@go run $(RUN_PATH)

run-local-docker-db:
	docker build -t go-exinity-task-postgress -f ./scripts/database/Dockerfile ./scripts/database/
	docker run --name go-exinity-task-postgress -p 5432:5432 -d go-exinity-task-postgress

run-docker-compose:
	docker-compose up

get-coverage:
	go test -cover -tags="!exclude_from_coverage" ./...

get-coverage-list:
	go tool cover -func=coverage.out

get-coverage-output:
	go test -tags="!exclude_from_coverage" -coverprofile=coverage.out ./...

get-coverage-output-html:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

generate-grpc:
	@protoc --go_out=./internal/app/api/proto  --go-grpc_out=./internal/app/api/proto ./internal/app/api/proto/candle.proto

document:
	echo "you can visit to localhost:6060 for auto generated documentation"
	godoc

init-swag:
	swag init -g ./cmd/api/main.go -o ./docs/swagger
