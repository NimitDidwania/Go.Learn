package handlers

import (
	"encoding/json"
	"learn/services/interfaces"
	"net/http"
)

type UserHandler struct {
	UserService interfaces.IUserService
}

// NewEchoHandler builds a new EchoHandler.
func NewUserHandler(userService interfaces.IUserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

// ServeHTTP handles an HTTP request to the /echo endpoint.
func (s *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := s.UserService.GetUser(0)
	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
