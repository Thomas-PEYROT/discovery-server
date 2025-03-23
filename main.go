package main

import (
	"discovery-server/routes"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {

	// Load dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("SERVER_PORT")

	http.HandleFunc("/register", routes.RegisterMicroservice)
	http.HandleFunc("/microservices", routes.GetAllMicroservices)
	fmt.Printf("Started server on port %v\n", port)
	http.ListenAndServe(":"+port, nil)
}
