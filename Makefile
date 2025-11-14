include .env

.DEFAULT_GOAL := help

GEN_SWAGGER_DIR := api/gen/swagger
SWAGGER_FILE := api/swagger/swagger.yml
POSTGRES_NAME := questions-postgres
POSTGRES_PORT := 1342
POSTGRES_USER := questions
POSTGRES_PASSWORD := 0000

.PHONY: help gen-swagger run lint cover postgres.start postgres.stop migrate.up migrate.down migrate.status tests swagger.start swagger.stop

help:
	@echo "Available targets:"
	@echo "  gen-swagger   - Generate swagger models"
	@echo "  run           - Run application"
	@echo "  lint          - Run linters"
	@echo "  cover         - Run tests with coverage"
	@echo "  postgres.start - Start PostgreSQL container"
	@echo "  postgres.stop  - Stop PostgreSQL container"
	@echo "  migrate.up    - Run migrations up"
	@echo "  migrate.down  - Run migrations down"
	@echo "  migrate.status - Show migration status"
	@echo "  tests         - Run tests"
	@echo "  swagger.start - Start swagger UI"
	@echo "  swagger.stop  - Stop swagger UI"

check-swagger:
	@test -f $(SWAGGER_FILE) || (echo "Error: swagger.yml not found" && exit 1)

gen-swagger:
	rm -rf $(GEN_SWAGGER_DIR)
	mkdir -p $(GEN_SWAGGER_DIR)
	docker run --rm --user $(shell id -u):$(shell id -g) -e GOPATH=$(go env GOPATH):/go \
	 	-v ${HOME}:${HOME} -w $(shell pwd) quay.io/goswagger/swagger:0.31.0 \
	 	generate model --spec=$(SWAGGER_FILE) --target=$(GEN_SWAGGER_DIR)
	go mod tidy

