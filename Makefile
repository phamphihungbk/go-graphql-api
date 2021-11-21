.PHONY: docker-up ## start docker container
up:
	docker-compose -f ./environment/development.yml --env-file .env up -d --remove-orphans

.PHONY: docker-build ## build docker image
build: copy-files
	docker-compose -f ./environment/development.yml --env-file .env build

.PHONY: docker-stop ## stop docker instance
down:
	docker-compose -f ./environment/development.yml down --remove-orphans

.PHONY: db-create
db-create:
	docker exec -it graphql-db sh -c "bash /docker-entrypoint-initdb.d/createDB.sh"

.PHONY: server-start
server-start:
	docker exec -it graphql-server sh -c "go run main.go"

.PHONY: server-generate
server-generate:
	docker exec -it graphql-server sh -c "go generate main.go"

.PHONY: server-build
server-build:
	docker exec -it graphql-server sh -c "go build main.go"

.PHONY: copy-files
copy-files:
	cp ./config/env.dev .env

.PHONY: db-list ## list all database tables
db-list:
	docker exec -it graphql-db sh -c "psql -l"

.PHONY: db-list ## list all database tables
go-get:
	docker exec -it graphql-db sh -c "go get -d -v ./..."

# postgres query
# \c graphql_db connect to db
# \dt list out all the tables
