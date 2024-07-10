package routes

import (
	"learn/pkg/http/handlers"
	"net/http"
)

// SetupRoutes sets up the application routes
func SetupRoutes(healthCheckHandler *handlers.HealthcheckHandler, userHandler *handlers.UserHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/api/healthcheck", healthCheckHandler)
	mux.Handle("/api/user", userHandler)

	// Add more routes here as needed
	return mux
}
