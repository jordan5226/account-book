/*******************************************
*	Migrate Database
*	Jordan, 2023
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
	fmt.Println("TODO: migrate")

	m, err := migrate.New("file://lib/pgdb/migration/migrations", dsn)

	if err != nil {
		log.Fatal(err)
	} else if err := m.Up(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("ENDDO: migrate")
}
