package routes

import (
	"learn/pkg/http/handlers"
	"net/http"
)

// SetupRoutes sets up the application routes
func SetupRoutes(healthCheckHandler *handlers.HealthcheckHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/healthcheck", healthCheckHandler)
	// Add more routes here as needed
	return mux
}
