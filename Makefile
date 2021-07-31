.PHONY: docker-up ## start docker container
up:
	docker-compose -f ./environment/development.yml up -d --remove-orphans

.PHONY: docker-build ## build docker image
build: copy-files
	docker-compose -f ./environment/development.yml build

.PHONY: docker-stop ## stop docker instance
down:
	docker-compose -f ./environment/development.yml down --remove-orphans

.PHONY: test ## run unit tests
test:
	@echo "mode: count" > coverage-all.out \
    @$(foreach pkg,$(PACKAGES), \
    	go test -p=1 -cover -covermode=count -coverprofile=coverage.out ${pkg}; \
    	tail -n +2 coverage.out >> coverage-all.out;)

.PHONY: test-cover
test-cover: test ## run unit tests and show test coverage information
	go tool cover -html=coverage-all.out

.PHONY: lint
lint: ## run golint on all Go package
	@golint $(PACKAGES)

.PHONY: db-stop
db-stop: ## stop the database server
	docker stop postgres

.PHONY: db-create
db-create:
	docker exec -it go-mysql sh -c "mysql -u root < /docker-entrypoint-initdb.d/createdb.sql"

.PHONY: copy-files
copy-files:
	cp ./environment/mysql/docker-entrypoint-initdb.d/createdb.sql.example ./environment/mysql/docker-entrypoint-initdb.d/createdb.sql
	cp ./config/.env.local .env
