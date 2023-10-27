package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/qwet700/go-practice/pkg/config"
	"github.com/qwet700/go-practice/pkg/model"
	utils "github.com/qwet700/go-practice/pkg/utilities"
)

var newItem model.Item

func GetItem(w http.ResponseWriter, r *http.Request) {
	newItem := model.GetAllItems()
	res, _ := json.Marshal(newItem)
	w.Header().Set("Content-Type", "app/json")
	w.WriteHeader(http.StatusOK) // send http respond
	w.Write(res)                 // get all item in database
}

func GetItemID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := vars["itemId"]
	ID, err := strconv.ParseInt(itemId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	itemDetails := model.GetItembyID(ID)
	res, _ := json.Marshal(itemDetails)
	w.Header().Set("Content-Type", "app/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// func GetItemID(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	itemId := vars["itemId"]
// 	ID, err := strconv.ParseInt(itemId, 0, 0)
// 	if err != nil {
// 		fmt.Println("error while parsing")
// 	}
// 	itemDetails, _ := model.GetItembyID(ID)
// 	res, _ := json.Marshal(itemDetails)
// 	w.Header().Set("Content-Type", "app/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

func CreItem(w http.ResponseWriter, r *http.Request) {
	CreateItem := &model.Item{}
	utils.ParseBody(r, CreateItem)
	a := CreateItem.CreateItem()
	res, _ := json.Marshal(a)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DelItemID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := vars["itemId"]
	ID, err := strconv.ParseInt(itemId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	item := model.DeleteItem(ID)
	res, _ := json.Marshal(item)
	w.Header().Set("Content-Type", "app/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdItemID(w http.ResponseWriter, r *http.Request) {
	var updateitem = &model.Item{}
	utils.ParseBody(r, updateitem)
	vars := mux.Vars(r)
	itemId := vars["itemId"]
	ID, err := strconv.ParseInt(itemId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	itemDetails := model.GetItembyID(ID) // db, itemDetails := model.GetItembyID(ID)
	if updateitem.Name != "" {
		itemDetails.Name = updateitem.Name
	}
	if updateitem.Sender != "" {
		itemDetails.Sender = updateitem.Sender
	}
	db := config.Getdb() //
	db.Save(&itemDetails)
	res, _ := json.Marshal(itemDetails)
	w.Header().Set("Content-Type", "app/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
