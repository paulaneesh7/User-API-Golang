package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/paulaneesh7/Users_API/routes"
)


func main() {
	fmt.Println("Server is starting...")

	r := routes.Routes()

	corsObj := handlers.AllowedOrigins([]string{"*"})
	corsHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(corsObj, corsHeaders, corsMethods)(r)))

	fmt.Println("Server running on port 8080")
}