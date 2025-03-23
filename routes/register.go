package routes

import (
	"encoding/json"
	"net/http"
)

type PostRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type PostResponse struct {
	Message string `json:"message"`
}

func RegisterMicroservice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method, should be POST", http.StatusMethodNotAllowed)
		return
	}

	var req PostRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response := PostResponse{Message: "ok"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
