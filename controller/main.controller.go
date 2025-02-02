package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/paulaneesh7/Users_API/helpers"
	"github.com/paulaneesh7/Users_API/model"
)

// POST
func CreateUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	helpers.CreateUser(user)
	json.NewEncoder(w).Encode("User created successfully✅")
}


// GET - Many
func GetUsersController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := helpers.GetAllUsers()
	json.NewEncoder(w).Encode(users)
}



// PUT
func UpdateUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	
	updatedUser := helpers.UpdateUserById(id, user)
	json.NewEncoder(w).Encode(updatedUser)

}


// DELETE
func DeleteUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	helpers.DeleteUserById(id)
	json.NewEncoder(w).Encode("User deleted successfully✅")
}