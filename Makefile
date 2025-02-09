.PHONY: tools openapi.build

BIN_DIR := $(shell pwd)/tools

oapi-codegen := go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.0

openapi.build:
	@echo "Building OpenAPI from TypeSpec"
	pnpm tsp.compile
	@echo "Building Hooks from Orval"
	pnpm orval.build
	@echo "Building API Handlers from OApi CodeGen"
	$(oapi-codegen) -package privateapi -o app/internal/adapter/server/oapi/private/openapi.gen.go ./doc/api/openapi/openapi.privateapi.yaml
	$(oapi-codegen) -package publicapi -o app/internal/adapter/server/oapi/public/openapi.gen.go ./doc/api/openapi/openapi.publicapi.yaml
