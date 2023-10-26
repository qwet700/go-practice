package routes

import (
	"github.com/gorilla/mux"
	"github.com/qwet700/go-practice/pkg/controller"
)

var ImpItem = func(router *mux.Router) {
	router.HandleFunc("/item/", controller.CreItem).Methods("POST")
	router.HandleFunc("/item/", controller.GetItem).Methods("GET")
	router.HandleFunc("/item/{id}", controller.GetItemID).Methods("GET")
	router.HandleFunc("/item/{id}", controller.UpdItemID).Methods("PUT")
	router.HandleFunc("/item/{id}", controller.DelItemID).Methods("DELETE")
}
