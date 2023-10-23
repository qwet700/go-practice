package item_routes

import (
	"github.com/gorilla/mux"
)

var ImpItem = func(router *mux.Router) {
	router.HandleFunc("/item/", controller.createItem).Methods("POST")
	router.HandleFunc("/item/", controller.getItem).Methods("GET")
	router.HandleFunc("/item/{id}", controller.getItemID).Methods("GET")

}
