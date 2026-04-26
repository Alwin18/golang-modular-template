package user

import (
	apperrors "github.com/Alwin18/golang-module-template/internal/shared/errors"
	"github.com/Alwin18/golang-module-template/internal/shared/utils"

	"gorm.io/gorm"
)

// Service implements user business logic.
type Service struct {
	repo *Repository
}

// NewService creates a new user Service.
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// List returns a paginated list of users.
func (s *Service) List(page, perPage int) ([]UserResponse, int64, int, error) {
	offset, totalPage := utils.Pagination(page, perPage, 0)

	users, total, err := s.repo.FindAll(offset, perPage)
	if err != nil {
		return nil, 0, 0, apperrors.ErrInternal
	}

	_, totalPage = utils.Pagination(page, perPage, total)

	result := make([]UserResponse, len(users))
	for i, u := range users {
		result[i] = toUserResponse(u.ID, u.Name, u.Email, u.Role, u.IsActive)
	}

	return result, total, totalPage, nil
}

// GetByID returns a single user.
func (s *Service) GetByID(id uint) (*UserResponse, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apperrors.ErrNotFound
		}
		return nil, apperrors.ErrInternal
	}

	resp := toUserResponse(user.ID, user.Name, user.Email, user.Role, user.IsActive)
	return &resp, nil
}

// Update modifies a user's profile.
func (s *Service) Update(id uint, req UpdateUserRequest) (*UserResponse, error) {

	if err := s.repo.Update(id, map[string]interface{}{"name": req.Name}); err != nil {
		return nil, apperrors.ErrInternal
	}

	return s.GetByID(id)
}

// Delete soft-deletes a user.
func (s *Service) Delete(id uint) error {
	if _, err := s.repo.FindByID(id); err != nil {
		return apperrors.ErrNotFound
	}
	return s.repo.Delete(id)
}

func toUserResponse(id uint, name, email, role string, isActive bool) UserResponse {
	return UserResponse{
		ID:       id,
		Name:     name,
		Email:    email,
		Role:     role,
		IsActive: isActive,
	}
}
