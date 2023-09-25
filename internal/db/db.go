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

type Product struct {
	Id    int64 `gorm:"primarykey"`
	Code  string
	Price int64
}

func Connect() (*gorm.DB, error) {
	dsn := fmt.Sprint("host=", host,
		" user=", user,
		" password=", password,
		" dbname=", dbname,
		" port=", port,
		` sslmode=disable
		TimeZone=Europe/Moscow`)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&Product{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
