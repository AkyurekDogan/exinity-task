# Makefile for Go project
API="api"
RUN_PATH="./cmd/api/main.go"

# Build the Go project
run:
	@go run $(RUN_PATH)

local-setup:
	@echo "Setting up local environment..."
	cp .env.dist .env
	@which yq >/dev/null 2>&1 || { \
		echo "📦 Installing yq..."; \
		brew install yq; \
	}
	@yq -i '.database.host = "localhost"' config.yml
	@docker-compose -f docker-compose.local.yml up -d
	@echo "Local environment is up and running, you can run the code in debug mode now!"
	@echo "Done!"

compose-up:
	@echo "Setting up local environment..."
	@which yq >/dev/null 2>&1 || { \
		echo "📦 Installing yq..."; \
		brew install yq; \
	}
	@yq -i '.database.host = "go-exinity-task-postgress"' config.yml
	@echo "Docker container is getting up..."
	@docker-compose -f docker-compose.yml up 
	@echo "Done!"

terraform-init:
	terraform init
	terraform apply

terraform-deploy-dev:
	terraform init
	terraform apply -var-file="environments/dev/terraform.tfvars"

terraform-deploy-staging:
	terraform init
	terraform apply -var-file="environments/staging/terraform.tfvars"

terraform-deploy-production:
	terraform init
	terraform apply -var-file="environments/prod/terraform.tfvars"

docker-compose:
	docker-compose up

generate-grpc:
	@protoc --go_out=./internal/app/proto  --go-grpc_out=./internal/app/proto ./internal/app/proto/candle.proto
