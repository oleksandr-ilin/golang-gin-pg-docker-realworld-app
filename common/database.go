package common

import (
	"fmt"
    "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func Init() *gorm.DB {
	dsn := "user=real-world-app password=123456 dbname=go-real-world-app-db port=5432 sslmode=disable"
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		fmt.Println("db err: ", err)
	}
	//db.DB().SetMaxIdleConns(10)
	//db.LogMode(true)
	DB = db
	return DB
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}
