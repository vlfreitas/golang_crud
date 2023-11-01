package infra

import (
	"fmt"

	"gorm.io/gorm"
)

const (
	DB_NAME = "golang_crud"
	DB_HOST = "localhost"
	DB_USER = "root"
	DB_PASS = "12345678"
	DB_PORT = "3306"
)

func Connect() *gorm.DB {

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME))
	if err != nil {
		fmt.Errorf("erro ao conectar com database")
	}

	return db
}
