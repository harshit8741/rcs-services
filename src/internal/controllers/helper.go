package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const (
	apiKey = "WAKlNGLQqaGIPLPq7kHM8QexG2wrATPv"
)

func respondWithJSON(w http.ResponseWriter, code int, response []byte, contentType string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func decodeJSONPayload(r *http.Request, v interface{}, w http.ResponseWriter) bool {
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		http.Error(w, "Invalid JSON Payload", http.StatusBadRequest)
		return false
	}
	return true
}

func sendRequestToExternalAPI(w http.ResponseWriter, url string, payload interface{}) {
	// Marshal the payload into JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Failed to prepare payload", http.StatusInternalServerError)
		return
	}

	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, "Error while creating request", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-apikey", apiKey)
	// fmt.Println("req processed", req)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error calling external API: %v", err)
		http.Error(w, "Failed to fetch data from external API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read and return the response
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response from external API", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, resp.StatusCode, responseBody, resp.Header.Get("Content-Type"))
}
