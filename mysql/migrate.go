package main

import (
	"github.com/yuuis/cat-api-go/adapter/datastore/cat"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

func main() {
	rootPassword := os.Getenv("MYSQL_ROOT_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	dbHost := os.Getenv("MYSQL_HOST")

	db, err := gorm.Open("mysql", "root:"+rootPassword+"@tcp("+dbHost+":3306)/"+dbName+"?charset=utf8mb4&parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	db.CreateTable(&datastoreCat.Cat{})
}
