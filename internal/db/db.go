package db

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func getDSN() string {
	dsn := fmt.Sprint("host=", os.Getenv("DB_HOST"),
		" user=", os.Getenv("POSTGRES_USER"),
		" password=", os.Getenv("POSTGRES_PASSWORD"),
		" dbname=", os.Getenv("DB_NAME"),
		" port=", os.Getenv("DB_PORT"),
		" sslmode=disable TimeZone=Europe/Moscow")
	return dsn
}

func Connect() (*gorm.DB, error) {
	log.Info(getDSN())
	db, err := gorm.Open(postgres.Open(getDSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&User{}, &RCEvent{}, &Pair{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
