package infra

import (
	"fmt"
	"log"
	"os"

	_ "time/tzdata"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBHandler struct {
	DB *gorm.DB
}

func NewDBHandler() *DBHandler {
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
	return &DBHandler{DB: db}
}

func (h *DBHandler) Find(dest interface{}, conds ...interface{}) (err error) {
	result := h.DB.Find(dest, conds)
	err = result.Error
	return
}

func (h *DBHandler) Create(value interface{}) (err error) {
	result := h.DB.Create(value)
	err = result.Error
	return
}
