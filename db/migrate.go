package db

import (
	"log"

	"github.com/yuuis/cat-api-go/adapter/datastore/mysql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("mysql", "root:password@/diffmvca?charset=utf8mb4&parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	db.CreateTable(&mysql.Item{})
}
