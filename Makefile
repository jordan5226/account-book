.PHONY: run migrateup migratedown

run:
	go run .

migrateup:
	go run lib/pgdb/migration/main.go up

migratedown:
	go run lib/pgdb/migration/main.go down

run: migrateup
