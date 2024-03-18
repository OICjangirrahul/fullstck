package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// RespondWithJSON marshals the provided data to JSON and writes it to the response writer
func RespondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	// Create the response object
	response := Response{
		Status: "success",
		Data:   data,
	}

	// Marshal the response object to JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding response to JSON", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json before writing to the response
	w.Header().Set("Content-Type", "application/json")

	// Respond with the JSON-encoded response
	w.WriteHeader(statusCode)
	w.Write(responseJSON)
}
