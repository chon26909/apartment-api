package services

import (
	"net/http"
	"time"

	"github.com/chon26909/apartment-api/models"
	"github.com/chon26909/apartment-api/repositories"
)

type IRoomService interface {
	GetAllRoom() (models.GetRoomsResponse, error)
}

type roomService struct {
	roomRepository repositories.IRoomRepository
}

func NewRoomService(roomRepository repositories.IRoomRepository) IRoomService {
	return &roomService{roomRepository: roomRepository}
}

func (r *roomService) GetAllRoom() (models.GetRoomsResponse, error) {

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	_ = client

	return models.GetRoomsResponse{}, nil
}
