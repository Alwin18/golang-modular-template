package auth

import (
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

func (r *Repository) Login(username, password string) (models.User, error) {
	var user models.User
	if err := r.db.Preload("Role").Where("username = ?", username).First(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *Repository) SaveRefreshToken(rt *models.RefreshToken) error {
	return r.db.Create(rt).Error
}

func (r *Repository) DeleteRefreshToken(token string) error {
	return r.db.Where("token = ?", token).Delete(&models.RefreshToken{}).Error
}
