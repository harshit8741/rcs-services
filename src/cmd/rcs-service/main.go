package main

import (
	"log"
	"net/http"
	"rcs-go-api/src/internal/controllers"

	"github.com/gorilla/mux"
	// "rcs-go-api/src/internal/messaging"
)

func main() {
	// router := messaging.SetupRouter()
	r := mux.NewRouter()
	r.HandleFunc("/api/send-message", controllers.FetchDataFromExternalAPI).Methods("POST")
	port := "8080"
	log.Printf("Server started on http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
