package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:password@tcp(localhost:3306)/bookstoreDB?charset=utf8&parseTime=True&loc=UTC")
	if err != nil {
		log.Printf("error while database connection %s: ", err)
	}

	// Test the connection
	if err := d.DB().Ping(); err != nil {
		log.Fatalf("Error while testing the database connection: %v", err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
