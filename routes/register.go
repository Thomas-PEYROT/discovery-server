package routes

import (
	"discovery-server/dto"
	"discovery-server/globals"
	"encoding/json"
	"github.com/google/uuid"
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
	if _, ok := globals.RegisteredMicroservices[body.Name]; ok {
		// We create the new instance of the microservice
		// TODO: Replace 5000 by a random port in a given range
		globals.RegisteredMicroservices[body.Name][id] = globals.MicroserviceInstance{Port: 5000}
	} else {
		// We register the new microservice
		globals.RegisteredMicroservices[body.Name] = make(map[string]globals.MicroserviceInstance)

		// We create the new instance of the microservice
		// TODO: Replace 5000 by a random port in a given range
		globals.RegisteredMicroservices[body.Name][id] = globals.MicroserviceInstance{Port: 5000}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.RegisterMicroserviceResponse{
		UUID: id,
		Port: 5000,
	})
}
