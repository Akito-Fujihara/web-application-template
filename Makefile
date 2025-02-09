.PHONY: openapi.build wire.build lint

BIN_DIR := $(shell pwd)/tools

oapi-codegen := go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
golangci-lint := go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest
wire := go run github.com/google/wire/cmd/wire@latest

openapi.build:
	@echo "Building OpenAPI from TypeSpec"
	pnpm tsp.compile
	@echo "Building Hooks from Orval"
	pnpm orval.build
	@echo "Building API Handlers from OApi CodeGen"
	$(oapi-codegen) -package privateapi -o app/adapter/server/oapi/privateapi/openapi.gen.go ./doc/api/openapi/openapi.privateapi.yaml
	$(oapi-codegen) -package publicapi -o app/adapter/server/oapi/publicapi/openapi.gen.go ./doc/api/openapi/openapi.publicapi.yaml

wire.build:
	@echo "Building Wire"
	$(wire) ./cmd/server/di/wire.go

lint:
	@echo "Running Lint"
	$(golangci-lint) run
	pnpm lint
	pnpm format
