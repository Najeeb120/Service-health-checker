package main

import (
	"log"

	"github.com/Najeeb120/Go-project/internal/checker"
)

func main() {

	go checker.StartHealthChecker(10)

	log.Println("Service Health Checker running on different ports")
	checker.ServeHealthStatus()
}
