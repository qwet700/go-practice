package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB // interact others db

func Connect() {
	a, err := gorm.Open("mysql", "root:root1@tcp(172.17.0.2:3306)/storage")
	if err != nil {
		panic(err)
	}
	db = a
}

func Getdb() *gorm.DB {
	return db
}
