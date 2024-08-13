package main

import (
	"anciledigital/configs"
	"anciledigital/controllers"
	"anciledigital/models"
	"net/http"
)

func main() {
	db := configs.GetDB()
	db.AutoMigrate(&models.User{})
	http.HandleFunc("/users", controllers.GetUsers)
	http.ListenAndServe(":8080", nil)
}
