package infra

import (
	"fmt"
	"log"
	"os"

	"github.com/ryoutaku/simple-chat/app/infra/adapter"

	_ "time/tzdata"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() *adapter.Database {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("DATABASE_NAME")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		user, password, host, port, dbName, "Asia%2FTokyo",
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	return &adapter.Database{DB: db}
}
