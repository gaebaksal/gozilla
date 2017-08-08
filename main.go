package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func main() {
	log.Fatal(http.ListenAndServe(":3000", mux.NewRouter()))
}

