GO_FILES = $(shell go list ./...)
GO_MODULE = assignment2
FLIP_DATABASE_URL :=  host=localhost user=indra password=pass1234 dbname=hacktiv port=5432 sslmode=disable

.PHONY: init
init:
	go mod tidy
	go mod vendor
	go mod verify

.PHONY: run
run: init
	@echo ${FLIP_DATABASE_URL}
	go run main.go

.PHONY: run-docker
run-docker:
	docker-compose up --build