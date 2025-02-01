package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/paulaneesh7/Users_API/routes"
)


func main() {
	fmt.Println("Server is starting...")

	r := routes.Routes()
	log.Fatal(http.ListenAndServe(":8080", r))

	fmt.Println("Server running on port 8080")
}