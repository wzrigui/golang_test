package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetGifts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var gifts []Gift
	DB.Find(&gifts)
	json.NewEncoder(w).Encode(gifts)
}

func GetGift(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var gift Gift
	DB.First(&gift, params["id"])
	json.NewEncoder(w).Encode(gift)
}

func CreateGift(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var gift Gift
	json.NewDecoder(r.Body).Decode(&gift)
	DB.Create(&gift)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(gift)
}

func UpdateGift(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var gift Gift
	DB.First(&gift, params["id"])
	json.NewDecoder(r.Body).Decode(&gift)
	DB.Save(&gift)
	json.NewEncoder(w).Encode(gift)
}
