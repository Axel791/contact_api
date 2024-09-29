package v1

import (
	"encoding/json"
	dto "github.com/Axel791/contact_api/internal/dto/v1"
	repo "github.com/Axel791/contact_api/internal/repositories/user"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type UserHandler struct {
	repo      repo.UserRepositoryInterface
	validator *validator.Validate
}

func NewUserDetailGetHandler(repo repo.UserRepositoryInterface) *UserHandler {
	return &UserHandler{
		repo:      repo,
		validator: validator.New(),
	}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	users, err := h.repo.GetAllUsers()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var userSummaries []dto.UserSummary
	for _, user := range users {
		userSummary := dto.UserSummary{
			ID:          int(user.ID),
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Age:         user.Age,
			Description: user.Description,
		}
		userSummaries = append(userSummaries, userSummary)
	}

	response := dto.UserList{
		Data: userSummaries,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
