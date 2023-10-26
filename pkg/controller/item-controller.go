package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/qwet700/go-practice/pkg/model"
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
	utils.parseBody(r, CreateItem)
	a := CreateItem.CreateItem()
	res, _ := json.Marshal(a)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
