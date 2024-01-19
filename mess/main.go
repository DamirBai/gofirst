package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestBody struct {
	Message string `json:"message"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

const expectedField = "message"
const successMessage = "Data successfully received"

func main() {
	http.HandleFunc("/json-endpoint", handleJSONRequest)
	fmt.Println("Server is listening on :8080...")
	http.ListenAndServe(":8080", nil)
}

func handleJSONRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var requestBody RequestBody

	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if requestBody.Message != "" {
		fmt.Println("Received message:", requestBody.Message)
		sendJSONResponse(w, http.StatusOK, successMessage)
	} else {
		sendJSONResponse(w, http.StatusBadRequest, "Invalid JSON message")
	}
}

func sendJSONResponse(w http.ResponseWriter, statusCode int, message string) {
	response := Response{
		Status:  "success",
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(response)
}