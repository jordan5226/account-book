package pgdb

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func conn() (err error) {
	dsn := os.Getenv("PG_URL")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return
}

func GetConnect() *gorm.DB {
	for i := 0; i < 3; i++ {
		if err := conn(); err != nil {
			fmt.Printf("[AccoutBook] pgdb.GetConnect - Connect Failed! ( Error: %v )\n", err.Error())
			time.Sleep(2 * time.Second)
		} else {
			return db
		}
	}

	return nil
}
