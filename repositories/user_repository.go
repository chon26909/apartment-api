package repositories

import (
	"github.com/chon26909/apartment-api/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserById(id int) (models.User, error)
	GetUserByEmail(email string) (user models.User, err error)
	GetAllUser(id string) (user models.User, err error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserById(id int) (user models.User, err error) {

	err = r.db.First(&user, id).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetUserByEmail(email string) (user models.User, err error) {

	err = r.db.Where(&models.User{Email: email}).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetAllUser(id string) (user models.User, err error) {

	err = r.db.First(&user, id).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
