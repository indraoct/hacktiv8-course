GO_FILES = $(shell go list ./...)
GO_MODULE = hacktiv8-course/final_assignment
export HACKTIV_DATABASE_STRING = host=localhost user=indra password=pass1234 dbname=mygram port=5432 sslmode=disable
export HACKTIV_PORT := 8090
export HACKTIV_PRIVATE_KEY := ./assets/keys/private_key.pem
export HACKTIV_PUBLIC_KEY := ./assets/keys/public_key.pem
export HACKTIV_SECRET_KEY := sfpas

.PHONY: init
init:
	go mod tidy
	go mod vendor
	go mod verify

.PHONY: run
run: init
	@echo ${HACKTIV_DATABASE_STRING}
	@echo ${HACKTIV_PORT}
	@echo ${HACKTIV_PRIVATE_KEY}
	@echo ${HACKTIV_PUBLIC_KEY}
	@echo ${HACKTIV_SECRET_KEY}
	go run main.go

.PHONY: run-docker
run-docker:
	docker-compose up --build