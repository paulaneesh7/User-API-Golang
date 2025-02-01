package controller

import (
	"encoding/json"
	"net/http"

	"github.com/paulaneesh7/Users_API/helpers"
	"github.com/paulaneesh7/Users_API/model"
)


func CreateUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	helpers.CreateUser(user)
	json.NewEncoder(w).Encode("User created successfully✅")
}


func GetUsersController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := helpers.GetAllUsers()
	json.NewEncoder(w).Encode(users)
}