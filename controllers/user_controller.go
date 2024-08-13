package controllers

import (
	"anciledigital/repositories"
	"encoding/json"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userRepo := repositories.NewUserRepository()

	users, err := userRepo.GetAllUsers()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	jsonResponse, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}
