package routes

import (
	"github.com/gorilla/mux"
	"github.com/paulaneesh7/Users_API/controller"
)

func Routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/user", controller.CreateUserController).Methods("POST")
	router.HandleFunc("/api/users", controller.GetUsersController).Methods("GET")
	router.HandleFunc("/api/user/{id}", controller.UpdateUserController).Methods("PUT")
	router.HandleFunc("/api/user/{id}", controller.DeleteUserController).Methods("DELETE")

	return router
}
