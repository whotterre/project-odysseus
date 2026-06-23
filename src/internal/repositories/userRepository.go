package repositories

import (
	"github.com/whotterre/odysseus/src/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(*models.User) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) CreateUser(user *models.User) (*models.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
