package main

import (
	"log"
	"net/http"
	"rcs-go-api/src/internal/controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	//plain text
	r.HandleFunc("/api/rcs/send-message", controllers.SimpleTextMessage).Methods("POST")

	// Standalone
	r.HandleFunc("/api/rcs/standalone/send-img-link-message", controllers.SimpleImgWithLinkMessage).Methods("POST")
	r.HandleFunc("/api/rcs/standalone/send-reply-message", controllers.SuggestedReplyMessage).Methods("POST")
	r.HandleFunc("/api/rcs/standalone/send-calender-event-message", controllers.SuggestedCalenderEventMessage).Methods("POST")
	r.HandleFunc("/api/rcs/standalone/send-location-message", controllers.SuggestedLocationMessage).Methods("POST")
	r.HandleFunc("/api/rcs/standalone/send-dailer-message", controllers.SuggestedDailerMessage).Methods("POST")

	//carousel
	port := "8080"

	log.Printf("Server started on http://localhost:%s", port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
