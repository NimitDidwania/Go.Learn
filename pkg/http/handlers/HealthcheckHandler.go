package handlers

import (
	"encoding/json"
	models "learn/Models"
	"learn/services/interfaces"
	"net/http"
)

type HealthcheckHandler struct {
}

// NewEchoHandler builds a new EchoHandler.
func NewHealthcheckHandler(userService interfaces.IUserService) *HealthcheckHandler {
	return &HealthcheckHandler{}
}

// ServeHTTP handles an HTTP request to the /echo endpoint.
func (s *HealthcheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := models.HealthcheckResponse{
		Status: "Healthy",
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
