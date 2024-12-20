package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// ExternalAPIResponse represents the expected response from the external API
type ExternalAPIResponse struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func FetchDataFromExternalAPI(w http.ResponseWriter, r *http.Request) {
	var e Payload
	err := json.NewDecoder(r.Body).Decode(&e)

	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	fmt.Println(e)

	externalAPIURL := "https://rcsapi-uat.jiocx.com/api/v1/sendMessage"

	jsonData, err := json.Marshal(e)

	if err != nil {
		http.Error(w, "Failed to prepare payload", http.StatusInternalServerError)
		return
	}
	req, err := http.NewRequest("POST", externalAPIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	// resp, err := http.Post(externalAPIURL, "application/json", bytes.NewBuffer(e))
	// resp, err := http.Post(externalAPIURL, "application/json", bytes.NewBuffer(e))

	// Add the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Add custom key-value pair in the header
	req.Header.Set("x-apikey", "WAKlNGLQqaGIPLPq7kHM8QexG2wrATPv")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Printf("Error calling external API: %v", err)
		http.Error(w, "Failed to fetch data from external API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)

	if err != nil {
		http.Error(w, "Failed to read response from external API", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.WriteHeader(resp.StatusCode)

	w.Write(responseBody)

}
