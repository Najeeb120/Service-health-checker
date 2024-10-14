package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate a 50/50 chance of being unhealthy
	if rand.Intn(2) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Service B is Unhealthy"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Service B is Healthy"))
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/health", healthHandler)
	log.Println("Service B running on :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
