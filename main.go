package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/baba", HandlerBaBa).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func HandlerBaBa(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("yetu"))
}