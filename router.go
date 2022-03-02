package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func initRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/employees", GetEmployees).Methods("GET")
	r.HandleFunc("/employees/{id}", GetEmployee).Methods("GET")
	r.HandleFunc("/employees", CreateEmployee).Methods("POST")
	r.HandleFunc("/employees/{id}", UpdateEmployee).Methods("PUT")

	r.HandleFunc("/gifts", GetGifts).Methods("GET")
	r.HandleFunc("/gifts/{id}", GetGift).Methods("GET")
	r.HandleFunc("/gifts", CreateGift).Methods("POST")
	r.HandleFunc("/gifts/{id}", UpdateGift).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080",
		handlers.CORS(handlers.AllowedHeaders([]string{
			"X-Requested-With",
			"Content-Type",
			"Authorization"}),
			handlers.AllowedMethods([]string{
				"GET",
				"POST",
				"PUT",
				"HEAD",
				"OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(r)))

}
