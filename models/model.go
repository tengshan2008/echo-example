package models

// HealthCheck with message
type HealthCheck struct {
	Message string `json:"message"`
}

// Cat with message
type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
