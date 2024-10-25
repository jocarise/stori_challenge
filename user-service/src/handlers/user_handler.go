package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"user-service/src/models"
	"user-service/src/services"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var userDto models.CreateUserDTO

	if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
		log.Printf("error parsing body %s", err)
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	user, err := h.service.CreateUser(&userDto)
	if err != nil {
		log.Printf("error saving user %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *UserHandler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var userDto models.AuthUserDTO

	if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	token, err := h.service.AuthUser(&userDto)
	if err != nil {
		http.Error(w, "authentication failed: "+err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"token": token}); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
