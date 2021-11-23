.PHONY: up
up:
	docker-compose -f ./environment/development.yml --env-file .env up -d --remove-orphans

.PHONY: build
build: files-copy
	docker-compose -f ./environment/development.yml --env-file .env build

.PHONY: down
down:
	docker-compose -f ./environment/development.yml down --remove-orphans

.PHONY: db-create
db-create:
	docker exec -it graphql-db sh -c "bash /docker-entrypoint-initdb.d/createDB.sh"

.PHONY: sv-start
sv-start:
	docker exec -it graphql-server sh -c "go run main.go"

.PHONY: sv-generate
sv-generate:
	docker exec -it graphql-server sh -c "go generate main.go"

.PHONY: sv-build
sv-build:
	docker exec -it graphql-server sh -c "go build main.go"

.PHONY: init-packages
init-packages:
	docker exec -it graphql-db sh -c "go get -d -v ./..."

.PHONY: files-copy
files-copy:
	cp ./config/env.dev .env

.PHONY: db-list
db-list:
	docker exec -it graphql-db sh -c "psql -l"

# postgres query
# \c graphql_db connect to db
# \dt list out all the tables
