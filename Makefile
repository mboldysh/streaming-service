NAME := streaming-service
GO := go
PROJECT ?= github.com/mboldysh/$(NAME)

all: clean fmt lint vet build

.PHONY: build
build: ## Builds a static executable
	@echo "+ $@"
	CGO_ENABLED=0 $(GO) build \
		-tags "$(BUILDTAGS) static_build" \
		${GO_LDFLAGS_STATIC} -o $(NAME) "$(PROJECT)/cmd/streaming-service"

.PHONY: fmt
fmt: ## Verifies all files have been `gofmt`ed
	@echo "+ $@"
	@gofmt -s -l -w . | grep -v vendor | tee /dev/stderr

.PHONY: lint
lint: ## Verifies `golint` passes
	@echo "+ $@"
	@golint ./... | grep -v vendor | tee /dev/stderr

.PHONY: vet
vet: ## Verifies `go vet` passes
	@echo "+ $@"
	@$(GO) vet $(shell $(GO) list ./... | grep -v vendor) | tee /dev/stderr

.PHONY: clean
clean:
	@echo "+ $@"
	$(RM) $(NAME)