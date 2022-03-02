package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees []Employee
	DB.Find(&employees)
	json.NewEncoder(w).Encode(employees)
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var employee Employee
	DB.First(&employee, params["id"])
	json.NewEncoder(w).Encode(employee)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	json.NewDecoder(r.Body).Decode(&employee)
	DB.Create(&employee)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(employee)

}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var employee Employee
	DB.First(&employee, params["id"])
	json.NewDecoder(r.Body).Decode(&employee)
	DB.Save(&employee)
	json.NewEncoder(w).Encode(employee)
}
