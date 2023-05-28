.PHONY: run migrate

run:
	go run .

migrate:
	go run lib/pgdb/migration/main.go

run: migrate
