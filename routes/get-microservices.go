package routes

import (
	"discovery-server/globals"
	"encoding/json"
	"net/http"
)

func GetAllMicroservices(w http.ResponseWriter, r *http.Request) {
	// Some checks
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method, should be GET", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(globals.RegisteredMicroservices)
}
