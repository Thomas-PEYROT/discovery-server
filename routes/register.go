package routes

import (
	"discovery-server/dto"
	"discovery-server/globals"
	"discovery-server/services"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func RegisterMicroservice(w http.ResponseWriter, r *http.Request) {
	// Some checks
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method, should be POST", http.StatusMethodNotAllowed)
		return
	}
	var body dto.RegisterMicroserviceBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id := uuid.New().String()

	// If an instance (or more) of this microservice has already been created
	if _, ok := globals.RegisteredMicroservices[body.Name]; !ok {
		// We register the new microservice
		globals.RegisteredMicroservices[body.Name] = make(map[string]globals.MicroserviceInstance)
	}

	// We create the new instance of the microservice
	globals.RegisteredMicroservices[body.Name][id] = globals.MicroserviceInstance{Port: services.GetNewPort()}
	log.Printf("Registered microservice \"%v\" on port %v (uuid: %v)", body.Name, globals.RegisteredMicroservices[body.Name][id].Port, id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.RegisterMicroserviceResponse{
		UUID: id,
		Port: globals.RegisteredMicroservices[body.Name][id].Port,
	})
}
