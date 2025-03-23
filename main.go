package main

import (
	"discovery-server/globals"
	"discovery-server/routes"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Load dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load SERVER_PORT
	port, err := strconv.ParseUint(os.Getenv("SERVER_PORT"), 10, 32)
	if err != nil {
		log.Fatal("Error parsing SERVER_PORT. Should be an uint32.")
	}
	globals.ServerPort = uint32(port)

	// Load MICROSERVICES_PORT_RANGE
	portRange := strings.Split(os.Getenv("MICROSERVICES_PORT_RANGE"), "-")
	if len(portRange) != 2 {
		log.Fatal("Error parsing MICROSERVICES_PORT_RANGE. Should be two uint32 separated by a \"-\" (for example 1000-2000).")
	}

	portRangeMin, err := strconv.ParseUint(portRange[0], 10, 32)
	if err != nil {
		log.Fatal("Error parsing MICROSERVICES_PORT_RANGE minimum value. Should be an uint32.")
	}
	globals.PortRangeMin = uint32(portRangeMin)

	portRangeMax, err := strconv.ParseUint(portRange[1], 10, 32)
	if err != nil {
		log.Fatal("Error parsing MICROSERVICES_PORT_RANGE minimum value. Should be an uint32.")
	}
	globals.PortRangeMax = uint32(portRangeMax)

	if globals.PortRangeMin > globals.PortRangeMax {
		log.Fatal("Error : Port range min should be greater than or equal to Port range max.")
	}

	// Start server and register endpoints
	http.HandleFunc("/register", routes.RegisterMicroservice)
	http.HandleFunc("/microservices", routes.GetAllMicroservices)
	fmt.Printf("Started server on port %v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
