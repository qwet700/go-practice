/*package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Object struct {
	ID       string  `json:"ID"`
	Name     string  `json:"name"`
	Describe string  `json:"describe"`
	Sender   *Sender `json:"sender"`
}

type Sender struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var objects []Object

func getObjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Typ", "app/json")
	json.NewEncoder(w).Encode(objects)
}

func delObject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Typ", "app/json")
	param := mux.Vars(r)
	for i, item := range objects {
		if item.ID == param["id"] {
			objects = append(objects[:i], objects[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(objects)

}

func getObject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Typ", "app/json")
	json.NewEncoder(w).Encode(objects)
	param := mux.Vars(r)
	for _, item := range objects {
		if item.ID == param["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func creObject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Typ", "app/json")
	var object Object
	_ = json.NewDecoder(r.Body).Decode(&object)
	object.ID = strconv.Itoa(rand.Intn(100000))
	objects = append(objects, object)
	json.NewEncoder(w).Encode(objects)
}

func updObject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Typ", "app/json")
	param := mux.Vars(r)
	for i, item := range objects {
		if item.ID == param["id"] {
			objects = append(objects[:i], objects[i+1:]...)
			var object Object
			_ = json.NewDecoder(r.Body).Decode(&object)
			object.ID = param["id"]
			objects = append(objects, object)
			return
		}
	}
	json.NewEncoder(w).Encode(objects)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/objects", getObjects).Methods("GET")
	r.HandleFunc("/objects/{id}", getObject).Methods("GET")
	r.HandleFunc("/objects", creObject).Methods("POST")
	r.HandleFunc("/objects/{id}", updObject).Methods("PUT")
	r.HandleFunc("/objects/{id}", delObject).Methods("DELETE")

	fmt.Printf("Starting the cargo on port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}*/