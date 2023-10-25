package item_routes

import (
	"github.com/gorilla/mux"
)

var ImpItem = func(router *mux.Router) {
	router.HandleFunc("/item/", controller.creItem).Methods("POST")
	router.HandleFunc("/item/", controller.getItem).Methods("GET")
	router.HandleFunc("/item/{id}", controller.getItemID).Methods("GET")
	router.HandleFunc("/item/{id}", controller.updItemID).Methods("PUT")
	router.HandleFunc("/item/{id}", controller.delItemID).Methods("DELETE")
}
