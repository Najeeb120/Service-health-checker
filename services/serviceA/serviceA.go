package main

import (
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Service A is Healthy"))
}

func main() {
	http.HandleFunc("/health", healthHandler)
	log.Println("Service A running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
