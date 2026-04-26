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

func (r *Repository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ? AND is_active = true", email).First(&user).Error
	return &user, err
}

func (r *Repository) FindUserByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *Repository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *Repository) SaveRefreshToken(token *models.RefreshToken) error {
	return r.db.Create(token).Error
}

func (r *Repository) FindRefreshToken(token string) (*models.RefreshToken, error) {
	var rt models.RefreshToken
	err := r.db.Where("token = ?", token).First(&rt).Error
	return &rt, err
}

func (r *Repository) DeleteRefreshToken(token string) error {
	return r.db.Where("token = ?", token).Delete(&models.RefreshToken{}).Error
}

func (r *Repository) DeleteUserRefreshTokens(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.RefreshToken{}).Error
}
