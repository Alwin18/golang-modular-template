package auth

import (
	"context"
	"errors"
	"time"

	"github.com/Alwin18/golang-module-template/internal/shared/cache"
	"github.com/Alwin18/golang-module-template/internal/shared/db/models"
	apperrors "github.com/Alwin18/golang-module-template/internal/shared/errors"
	"github.com/Alwin18/golang-module-template/internal/shared/security"

	"gorm.io/gorm"
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

// Register creates a new user account.
func (s *Service) Register(req RegisterRequest) (*TokenResponse, error) {

	// Check duplicate email
	_, err := s.repo.FindUserByEmail(req.Email)
	if err == nil {
		return nil, apperrors.ErrConflict
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperrors.ErrInternal
	}

	hashed, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, apperrors.ErrInternal
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashed,
		Role:     "user",
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, apperrors.ErrInternal
	}

	return s.generateTokens(user)
}

// Login authenticates a user.
func (s *Service) Login(req LoginRequest) (*TokenResponse, error) {
	user, err := s.repo.FindUserByEmail(req.Email)
	if err != nil {
		return nil, apperrors.ErrUnauthorized
	}

	if !security.CheckPassword(req.Password, user.Password) {
		return nil, apperrors.ErrUnauthorized
	}

	return s.generateTokens(user)
}

// Refresh rotates tokens using a valid refresh token.
func (s *Service) Refresh(req RefreshRequest) (*TokenResponse, error) {
	claims, err := s.jwtManager.ParseToken(req.RefreshToken)
	if err != nil {
		return nil, apperrors.ErrUnauthorized
	}

	rt, err := s.repo.FindRefreshToken(req.RefreshToken)
	if err != nil || rt.ExpiresAt.Before(time.Now()) {
		return nil, apperrors.ErrUnauthorized
	}

	user, err := s.repo.FindUserByID(claims.UserID)
	if err != nil {
		return nil, apperrors.ErrUnauthorized
	}

	// Rotate: delete old token
	_ = s.repo.DeleteRefreshToken(req.RefreshToken)

	return s.generateTokens(user)
}

// Logout invalidates the user's refresh token.
func (s *Service) Logout(userID uint, refreshToken string) error {
	_ = s.repo.DeleteRefreshToken(refreshToken)
	return nil
}

// Me returns the user profile.
func (s *Service) Me(userID uint) (*MeResponse, error) {
	user, err := s.repo.FindUserByID(userID)
	if err != nil {
		return nil, apperrors.ErrNotFound
	}

	return &MeResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}

// generateTokens creates and persists access + refresh tokens.
func (s *Service) generateTokens(user *models.User) (*TokenResponse, error) {
	accessToken, err := s.jwtManager.GenerateAccessToken(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, apperrors.ErrInternal
	}

	refreshToken, expiresAt, err := s.jwtManager.GenerateRefreshToken(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, apperrors.ErrInternal
	}

	rt := &models.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: expiresAt,
	}
	if err := s.repo.SaveRefreshToken(rt); err != nil {
		return nil, apperrors.ErrInternal
	}

	// Cache user info
	_ = s.cache.Set(context.Background(), cache.UserKey(user.ID), user.Email, 15*time.Minute)

	return &TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
	}, nil
}
