package http

import (
	"encoding/json"
	"net/http"
	"strings"

	"fastapi2/hexagonal/application/usecase"
)

type UserHandler struct {
	createUser *usecase.CreateUserUseCase
	findUser   *usecase.FindUserUseCase
}

func NewUserHandler(
	createUser *usecase.CreateUserUseCase,
	findUser *usecase.FindUserUseCase,
) *UserHandler {
	return &UserHandler{
		createUser: createUser,
		findUser:   findUser,
	}
}

type createUserRequest struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type userResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request createUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	user, err := h.createUser.Execute(r.Context(), usecase.CreateUserInput{
		ID:    request.ID,
		Name:  request.Name,
		Email: request.Email,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := userResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/users/")
	if id == "" || id == r.URL.Path {
		http.Error(w, "missing user id", http.StatusBadRequest)
		return
	}

	user, err := h.findUser.Execute(r.Context(), usecase.FindUserInput{
		ID: id,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response := userResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}