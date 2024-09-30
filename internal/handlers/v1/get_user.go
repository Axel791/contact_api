package v1

import (
	"encoding/json"
	dto "github.com/Axel791/contact_api/internal/dto/v1"
	repo "github.com/Axel791/contact_api/internal/repositories/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

type UserDetailHandler struct {
	repo      repo.IUserRepository
	validator *validator.Validate
}

func NewUserDetailHandler(repo repo.IUserRepository) *UserDetailHandler {
	return &UserDetailHandler{
		repo:      repo,
		validator: validator.New(),
	}
}

func (h *UserDetailHandler) GetUserDetail(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "userId")
	userId64, err := strconv.ParseUint(userIdStr, 10, 64)

	userId := uint(userId64)

	user, err := h.repo.GetUserByID(userId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response := dto.UserSummary{
		ID:          int(user.ID),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Age:         user.Age,
		Description: user.Description,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
