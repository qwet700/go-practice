package model

import (
	"github.com/jinzhu/gorm"
	"github.com/qwet700/go-practice/pkg/config"
)

var db *gorm.DB

type Item struct {
	gorm.Model
	Name   string  `gorm:""json:"name"`
	Sender string  `json:"sender"`
	Price  float64 `json:"price"`
}

func init() {
	config.Connect()
	db = config.Getdb()
	db.AutoMigrate(&Item{})
}

func (b *Item) CreateItem() *Item {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllItems() []Item { // [] return a list of items
	var Items []Item
	db.Find(&Items)
	return Items
}

// func GetItembyID(id int64) (*Item, *gorm.DB) { // return the item equal to id and all db
// 	var getItem Item
// 	db := db.Where("id=?", id).Find(&getItem) // "id=?" where id == routes {id}
// 	return &getItem, db
// }

func GetItembyID(id int64) *Item {
	var getItem Item
	db.Where("id=?", id).Find(&getItem)
	return &getItem
}

func DeleteItem(id int64) Item {
	var item Item
	db.Where("id=?", id).Delete(item)
	return item
}
