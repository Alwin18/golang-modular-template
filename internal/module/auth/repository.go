package auth

import (
	"context"

	"github.com/Alwin18/golang-module-template/internal/shared/constants"
	"github.com/Alwin18/golang-module-template/internal/shared/db/models"
	"gorm.io/gorm"
)

// Repository handles auth data access.
type Repository struct {
	db *gorm.DB
}

// NewRepository creates a new auth Repository.
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Login(ctx context.Context, username string) (models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).Preload("Role").Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.User{}, constants.ErrAccountNotFound
		}
		return models.User{}, constants.ErrInternalServer
	}

	return user, nil
}

func (r *Repository) SaveRefreshToken(rt *models.RefreshToken) error {
	if err := r.db.Create(rt).Error; err != nil {
		return constants.ErrInternalServer
	}

	return nil
}

func (r *Repository) DeleteRefreshToken(token string) error {
	if err := r.db.Where("token = ?", token).Delete(&models.RefreshToken{}).Error; err != nil {
		return constants.ErrInternalServer
	}

	return nil
}
