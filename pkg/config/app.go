package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB // interact others db

func Connect() {
	// start docker mysql
	// sudo docker run --name mysql-db -e MYSQL_ROOT_PASSWORD=root1 -p 127.0.0.1:3306:3306 -d mysql
	//
	// connect sql
	// mysql -h 127.0.0.1 -P 3306 -uroot -proot1
	a, err := gorm.Open("mysql", "root:root1@tcp(172.17.0.2:3306)/storage?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = a
}

func Getdb() *gorm.DB {
	return db
}
