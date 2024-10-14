package checker

// Service represents a microservice to check
type Service struct {
	Name string
	URL  string
}

// List of services to monitor
var services = []Service{
	{"Service A", "http://localhost:8081/health"},
	{"Service B", "http://localhost:8082/health"},
	{"Service C", "http://localhost:8083/health"},
}
