.DEFAULT_GOAL := swagger
# Image URL to use all building/pushing image targets
IMG ?= quay.io/praveen4g0/url-shortner:latest

install_swagger:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger:
	@echo Ensure you have the swagger CLI or this command will fail.
	@echo You can install the swagger CLI with: go get -u github.com/go-swagger/go-swagger/cmd/swagger
	@echo ....

	swagger generate spec -o ./swagger.json --scan-models
fmt: ## Run go fmt against code.
	go fmt ./...

vet: ## Run go vet against code.
	go vet ./...

##@ Build

build: fmt vet ## Build manager binary.
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o url-shortner .

run: build ## Run a controller from your host.
	./url-shortner

docker-build: ## Build docker image with the manager.
	docker build -t ${IMG} .

docker-push: ## Push docker image with the manager.
	docker push ${IMG}		