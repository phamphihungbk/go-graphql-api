#!make
##@ [Docker] Build / Infrastructure
PACKAGE ?= example.com/mypackage

.PHONY: up
up: ## start docker instances
	docker-compose -f ./environment/development.yml --env-file .env up -d --remove-orphans

.PHONY: build
build: ## build docker images
	cp ./config/env.dev .env
	docker-compose -f ./environment/development.yml --env-file .env build

.PHONY: down
down: ## stop docker instances
	docker-compose -f ./environment/development.yml down --remove-orphans

.PHONY: db-create
db-create: ## create database on postgres instance
	docker exec -it graphql-db sh -c "bash /docker-entrypoint-initdb.d/createDB.sh"

.PHONY: sv-start
sv-start: ## start graphql api server
	docker exec -it graphql-server sh -c "go run main.go"

.PHONY: graphql-generate
graphql-generate: ## generate graphql schema from gqlgen config
	docker exec -it graphql-server sh -c "go run github.com/99designs/gqlgen generate"

.PHONY: sv-wire
sv-wire: ## generate wire
	docker exec -it graphql-server sh -c "go generate ./..."

.PHONY: sv-generate
sv-generate: ## generate google wire, graphql schema
	docker exec -it graphql-server sh -c "go generate main.go"

.PHONY: sv-build
sv-build: ## build production
	docker exec -it graphql-server sh -c "go build main.go"

.PHONY: package-install
package-install: ## install packages
	docker exec -it graphql-server sh -c "go get -d -v ./..."

.PHONY: package-upgrade
package-upgrade: ## upgrade packages with PACKAGE=
	docker exec -it graphql-server sh -c "go get -u ${@:1}"

.PHONY: lint
lint: ## run lint
	docker exec -it graphql-server sh -c "golangci-lint run --fix --fast"

help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-27s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

# postgres query
# \c graphql_db connect to db
# \dt list out all the tables
# docker exec -it graphql-db sh -c "psql -l"

# wire
# run these commands when get issue with @should not have @versio
# go clean -modcache
# go mod tidy
