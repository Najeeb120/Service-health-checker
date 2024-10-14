package main

import (
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Service C is Unhealthy"))
}

func main() {
	http.HandleFunc("/health", healthHandler)
	log.Println("Service C running on :8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
