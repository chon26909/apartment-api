package repositories

import (
	"github.com/chon26909/apartment-api/models"
	"gorm.io/gorm"
)

type IRoomRepository interface {
	GetAllRoom() (models.Room, error)
}

type roomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) IRoomRepository {
	return &roomRepository{db: db}
}

func (r *roomRepository) GetAllRoom() (models.Room, error) {

	return models.Room{}, nil
}
