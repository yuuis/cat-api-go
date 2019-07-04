package infrastructure

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
)

func OpenDB() (*gorm.DB, error) {
	rootPassword := os.Getenv("MYSQL_ROOT_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	dbHost := os.Getenv("MYSQL_HOST")

	return gorm.Open("mysql", "root:" + rootPassword + "@tcp("+dbHost+":3306)/" + dbName + "?charset=utf8mb4&parseTime=true")
}
