/*******************************************
*	Migrate Database
*	Jordan, 05/2023
********************************************/
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/joho/godotenv/autoload"
)

var dsn string

func init() {
	dsn = os.Getenv("PG_URL")
}

func main() {
	fmt.Printf("TODO: migrate %s\n", os.Args[1])

	path := os.Getenv("MIGRATE_PATH")
	m, err := migrate.New(path, dsn)

	if err != nil {
		log.Fatal(err)
	} else {
		switch os.Args[1] {
		case "up":
			if err := m.Up(); err != nil {
				if err != migrate.ErrNoChange {
					log.Fatal(err)
				}
			}
		case "down":
			if err := m.Down(); err != nil {
				if err != migrate.ErrNoChange {
					log.Fatal(err)
				}
			}
		}
	}

	fmt.Printf("ENDDO: migrate %s\n", os.Args[1])
}
