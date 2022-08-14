package handler

import (
	"encoding/json"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	message := map[string]string{
		"message": "Hello!",
	}

	json.NewEncoder(w).Encode(message)
}
