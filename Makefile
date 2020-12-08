GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=k8s-secret-auditor.exe
BINARY_OUTPUT_FOLDER_PATH=_bin
SERVICE_ENTRYPOINT=cmd/k8s-secret-auditor/main.go

.DEFAULT_GOAL = run

define printinfo
	@echo "\033[1;34m[!] $(1)\033[0m"
endef
define printsuccess
 @echo "\033[32m[+] $(1)\033[0m"
endef

# Basic operations
.PHONY: lint
lint:
	@GOPATH=${PWD}/.gopath golint ./...

.PHONY: build
build:
	$(eval BINARY_OUTPUT_FILE_PATH=$(BINARY_OUTPUT_FOLDER_PATH)/$(BINARY_NAME))
	$(call printinfo,Starting $(BINARY_NAME) build ...)
	@GOPATH=${PWD}/.gopath CGO_ENABLED=0 go build -gcflags="-m" -ldflags '-s -w' -o $(BINARY_OUTPUT_FILE_PATH) $(SERVICE_ENTRYPOINT)
	$(call printsuccess,Build completed successfully !)
	$(call printinfo,Binary available here : $(BINARY_OUTPUT_FILE_PATH))

.PHONY: run
run:
	@GOPATH=${PWD}/.gopath go run $(SERVICE_ENTRYPOINT)

