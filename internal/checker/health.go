package checker

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

var mu sync.Mutex
var serviceStatuses = make(map[string]string)

// checkHealth checks the health of a single service
func checkHealth(service Service, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(service.URL)
	mu.Lock()
	defer mu.Unlock()

	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Service %s is Unhealthy\n", service.Name)
		serviceStatuses[service.Name] = "Unhealthy"
	} else {
		log.Printf("Service %s is Healthy\n", service.Name)
		serviceStatuses[service.Name] = "Healthy"
	}
}

// StartHealthChecker initiates periodic health checking
func StartHealthChecker(interval int) {
	for {
		var wg sync.WaitGroup
		for _, service := range services {
			wg.Add(1)
			go checkHealth(service, &wg)
		}
		wg.Wait()
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

// ServeHealthStatus exposes the health status via HTTP
func ServeHealthStatus() {
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(serviceStatuses)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
