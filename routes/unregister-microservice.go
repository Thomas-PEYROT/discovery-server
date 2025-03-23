package routes

import (
	"discovery-server/dto"
	"discovery-server/exceptions"
	"discovery-server/globals"
	"encoding/json"
	"log"
	"net/http"
)

func UnregisterMicroservice(w http.ResponseWriter, r *http.Request) {
	// Some checks
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method, should be POST", http.StatusMethodNotAllowed)
		return
	}
	var body dto.UnregisterMicroserviceBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// If an instance (or more) of this microservice has already been created
	found := false
	for microservice, _ := range globals.RegisteredMicroservices {
		if instance, ok := globals.RegisteredMicroservices[microservice][body.UUID]; ok {
			// We remove the microservice
			delete(globals.RegisteredMicroservices[microservice], body.UUID)
			log.Printf("Unegistered microservice \"%v\" on port %v (uuid: %v)", microservice, instance.Port, body.UUID)
			found = true
		}
	}

	if !found {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(exceptions.HttpException{Message: "No microservice found for these parameters.", StatusCode: http.StatusBadRequest})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.UnregisterMicroserviceResponse{Message: "Successfully unregistered microservice."})
}
