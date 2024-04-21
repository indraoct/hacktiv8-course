GO_FILES = $(shell go list ./...)
GO_MODULE = hacktiv8-course/assignment2
export HACKTIV_DATABASE_STRING = host=localhost user=indra password=pass1234 dbname=hacktiv port=5432 sslmode=disable
export HACKTIV_PORT := 8090

.PHONY: init
init:
	go mod tidy
	go mod vendor
	go mod verify

.PHONY: run
run: init
	@echo ${HACKTIV_DATABASE_STRING}
	@echo ${HACKTIV_PORT}
	go run main.go

.PHONY: run-docker
run-docker:
	docker-compose up --build