BINARY_NAME=k8s-secret-auditor.exe
BINARY_OUTPUT_FOLDER_PATH=bin
SERVICE_ENTRYPOINT=cmd/k8s-secret-auditor/main.go

.DEFAULT_GOAL = run

.PHONY: lint
lint:
	@GOPATH=${PWD}/.gopath golint ./...

.PHONY: build
build:
	@GOPATH=${PWD}/.gopath CGO_ENABLED=0 go build -gcflags="-m" -ldflags '-s -w' -o bin/$(BINARY_NAME) $(SERVICE_ENTRYPOINT)

.PHONY: run
run:
	@GOPATH=${PWD}/.gopath go run $(SERVICE_ENTRYPOINT)

.PHONY: test
test:
	@GOPATH=${PWD}/.gopath go test ./...