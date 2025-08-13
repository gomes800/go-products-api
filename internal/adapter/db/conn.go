package bd

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupDB() *gorm.DB {
	dsn := "host=localhost user=admin password=12345 dbname=products_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection error:", err)
	}
	fmt.Println("Connected!")
	return db
}
