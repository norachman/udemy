package db

import (
	"errors"
	"day-1/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Creds struct{
	HostName 	string
	UserName	string
	Password	string
	DBName		string
	Port		string
}

var db *gorm.DB

func Connection() error {
	var err error

	postgre := "postgres://postgres:12januari@localhost:5432/course"
	
	dsn := postgre

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return errors.New("can't connect to database")
	}
	

	e := db.AutoMigrate(&model.User{})
	if e != nil {
		return e
	}
	return nil
}

func SetUpDBCOnnection(DB *gorm.DB) {
	db = DB
}

func GetConnectionDB() *gorm.DB {
	return db
}