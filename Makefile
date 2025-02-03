.PHONY: openapi.build

openapi.build:
	@echo "Building OpenAPI from TypeSpec"
	pnpm tsp.compile
