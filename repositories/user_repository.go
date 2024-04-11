package repositories

import (
	"github.com/chon26909/apartment-api/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserById(id string) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserById(id string) (models.User, error) {
	return models.User{}, nil
}
