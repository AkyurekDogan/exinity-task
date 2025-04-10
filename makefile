# Makefile for Go project
API="api"
RUN_PATH="./cmd/api/main.go"

# Tooling 
# --------------------------------------------------------------
generate-grpc:
	@protoc --go_out=./internal/app/proto  --go-grpc_out=./internal/app/proto ./internal/app/proto/candle.proto

# Local run setup
# --------------------------------------------------------------
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

# Docker Compose setup 
# --------------------------------------------------------------
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

# Kubernetes 
# --------------------------------------------------------------
kubernetes-setup:
	@echo "Setting up k8s namespace for postgres ..."
	@kubectl apply -f ./deploy/k8s/namespace.yml
	@echo "Setting up k8s secrets..."
	@kubectl apply -f ./deploy/k8s/postgres/secret.yml  
	@echo "Setting up k8s configmap..."
	@kubectl apply -f ./deploy/k8s/postgres/configmap.yml  
	@echo "Setting up k8s deployments..."
	@kubectl apply -f ./deploy/k8s/postgres/deployment.yml
	@echo "Setting up k8s services..."
	@kubectl apply -f ./deploy/k8s/postgres/service.yml
	@echo "Postgres is done!"
	
	@echo "Setting up k8s for gRPC Api ..."
	@echo "Setting up k8s secrets..."
	@kubectl apply -f ./deploy/k8s/api/secret.yml  
	@echo "Setting up k8s configmap..."
	@kubectl apply -f ./deploy/k8s/api/configmap.yml  
	@echo "Setting up k8s deployments..."
	@kubectl apply -f ./deploy/k8s/api/deployment.yml
	@echo "Setting up k8s services..."
	@kubectl apply -f ./deploy/k8s/api/service.yml
	@echo "API is done!"

connect-kuberneter-postgres:
	@echo "Routing the local connection to postgres pod..."
	@kubectl port-forward svc/postgres 5432:5432 -n exinity-task
	@echo "Done! You can connect via localhost:5432 now!"

# Terraform 
# --------------------------------------------------------------
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