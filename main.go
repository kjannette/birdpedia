package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", handler).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func handler(w http.ResponseWriter, router *http.Request) {
	fmt.Fprintf(w, "Hello Whirrrld EAT IT")
}
