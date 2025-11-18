include .env

.DEFAULT_GOAL := help

GEN_SWAGGER_DIR := api/gen/swagger
SWAGGER_FILE := api/swagger/swagger.yml
POSTGRES_NAME := questions
POSTGRES_PORT := 5434
POSTGRES_USER := questions
POSTGRES_PASSWORD := 0000

.PHONY: help gen-swagger run lint cover postgres.start postgres.stop migrate.up migrate.down migrate.status tests swagger.start swagger.stop

help:
	@echo "Available targets:"
	@echo "  gen-swagger   - Generate swagger models"
	@echo "  run           - Run application"
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

run:
	go run cmd/app/main.go -cfg configs/local.json

cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

postgres.start:
	if [ ! "$(shell docker ps -q -f name=$(POSTGRES_NAME))" ]; then \
        if [ "$(shell docker ps -aq -f status=exited -f name=$(POSTGRES_NAME))" ]; then \
            docker rm $(POSTGRES_NAME); \
        fi; \
		docker run --restart unless-stopped -d -p $(POSTGRES_PORT):5432 \
            -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) --name $(POSTGRES_NAME) postgres:13; \
        sleep 5; \
    fi;
	-docker exec $(POSTGRES_NAME) psql -U postgres -c "create user $(POSTGRES_USER) password '$(POSTGRES_PASSWORD)'"
	-docker exec $(POSTGRES_NAME) psql -U postgres -c "create database $(POSTGRES_USER)"
	-docker exec $(POSTGRES_NAME) psql -U postgres -c "grant all privileges on database $(POSTGRES_USER) to $(POSTGRES_USER)"

postgres.stop:
	docker stop $(POSTGRES_NAME)
	docker rm $(POSTGRES_NAME)

migrate.up:
	goose -dir migrations postgres "host=localhost port=$(POSTGRES_PORT) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_USER) sslmode=disable" up

migrate.down:
	goose -dir migrations postgres "host=localhost port=$(POSTGRES_PORT) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_USER) sslmode=disable" down

migrate.status:
	goose -dir migrations postgres "host=localhost port=$(POSTGRES_PORT) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_USER) sslmode=disable" status

tests:
	go test -v ./...

swagger.start:
	docker run --rm -d \
		--platform linux/amd64 \
		--name reviews-swagger \
		-p 9805:8080 \
		-e SWAGGER_JSON=/swagger.yaml \
		-v $(shell pwd)/api/swagger/swagger.yaml:/swagger.yaml \
		swaggerapi/swagger-ui

swagger.stop:
	docker stop reviews-swagger
	docker rm reviews-swagger