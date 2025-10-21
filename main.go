package main

import (
	"down_detective/routes"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// loading env and extracting port 
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")

	// making a router and passing it in app
	r := mux.NewRouter()
	routes.RegisterAppRoutes(r)

	// server
	fmt.Printf("Server is running on PORT %s\n", port)
	http.ListenAndServe(":"+port, r)
}
