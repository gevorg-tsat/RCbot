package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
	port     = 5432
)

const dsn = fmt.Sprintf("host=", host,
	"user=", user,
	"password=", password,
	"dbname=", dbname,
	"port=", port,
	`sslmode=disable 
TimeZone=Asia/Shanghai`)

func main() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
