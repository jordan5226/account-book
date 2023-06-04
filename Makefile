.PHONY: run rundb createdb dropdb migrateup migratedown test

run:
	go run .

rundb:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15.3-alpine3.18

createdb:
	docker exec -it postgres createdb --username=root --owner=root acctbook

dropdb:
	docker exec -it postgres dropdb acctbook

migrateup:
	go run lib/pgdb/migration/main.go up

migratedown:
	go run lib/pgdb/migration/main.go down

test:
	go test -v -cover ./...

run: migrateup
