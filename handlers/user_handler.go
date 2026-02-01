package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

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
		// Check if an "id" query parameter exists
		if r.URL.Query().Get("id") != "" {
			h.getUserByID(w, r)
		} else {
			h.getUsers(w, r)
		}
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

func (h *UserHandler) getUserByID(w http.ResponseWriter, r *http.Request) {
	// 1. Extract the ID from the URL query string
	idStr := r.URL.Query().Get("id")

	// 2. Convert the string ID to an integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID format", http.StatusBadRequest)
		return
	}

	// 3. Call the storage layer
	user, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		// If storage returns "user not found", send a 404
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// 4. Return the user as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
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
