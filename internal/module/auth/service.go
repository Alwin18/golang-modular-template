package auth

import (
	"context"
	"time"

	"github.com/Alwin18/golang-module-template/internal/shared/cache"
	"github.com/Alwin18/golang-module-template/internal/shared/constants"
	"github.com/Alwin18/golang-module-template/internal/shared/db/models"
	"github.com/Alwin18/golang-module-template/internal/shared/security"
)

// Service implements auth business logic.
type Service struct {
	repo       *Repository
	jwtManager *security.JWTManager
	cache      *cache.Cache
}

// NewService creates a new auth Service.
func NewService(repo *Repository, jwtManager *security.JWTManager, cacheClient *cache.Cache) *Service {
	return &Service{repo: repo, jwtManager: jwtManager, cache: cacheClient}
}

// Login authenticates a user.
func (s *Service) Login(req LoginRequest) (LoginResponse, error) {
	user, err := s.repo.Login(req.Username, req.Password)
	if err != nil {
		return LoginResponse{}, err
	}

	if !security.CheckPassword(req.Password, user.Password) {
		return LoginResponse{}, constants.ErrInvalidPassword
	}

	accessToken, err := s.jwtManager.GenerateAccessToken(user.ID, user.Email, user.Role.Name)
	if err != nil {
		return LoginResponse{}, constants.ErrInternalServer
	}

	refreshToken, expiresAt, err := s.jwtManager.GenerateRefreshToken(user.ID, user.Email, user.Role.Name)
	if err != nil {
		return LoginResponse{}, constants.ErrInternalServer
	}

	if err := s.repo.SaveRefreshToken(&models.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: expiresAt,
	}); err != nil {
		return LoginResponse{}, constants.ErrInternalServer
	}

	// Cache user info
	_ = s.cache.Set(context.Background(), cache.UserKey(user.ID), user.Username, 15*time.Minute)

	return LoginResponse{
		User: User{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Status:   user.Status,
			Role: Role{
				ID:   user.RoleID,
				Name: user.Role.Name,
			},
		},
		Token:        accessToken,
		RefreshToken: refreshToken,
	}, nil
}
