package handlers

import (
	"encoding/json"
	"net/http"

	"go-http-server/models"
	"go-http-server/storage"
)

type UserHandler struct {
	repo *storage.UserRepository
}

func NewUserHandler(repo *storage.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		h.getUsers(w, r)

	case http.MethodPost:
		h.createUser(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.repo.GetAll(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var input models.User

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if input.Name == "" || input.Email == "" {
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}

	user, err := h.repo.Create(r.Context(), input)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
