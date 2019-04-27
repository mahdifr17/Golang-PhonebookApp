package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/mahdifr17/phonebook-v2/controllers"
)

func main() {
	/* Routing */
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.GetAllRecord).Methods("GET")
	router.HandleFunc("/{id}", controllers.GetRecord).Methods("GET")
	router.HandleFunc("/", controllers.AddRecord).Methods("POST")
	router.HandleFunc("/{id}", controllers.UpdateRecord).Methods("PUT")
	router.HandleFunc("/{id}", controllers.DeleteRecord).Methods("DELETE")

	/* Start web server */
	http.ListenAndServe(":8080", router)
}
