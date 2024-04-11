package handlers

import (
	"fmt"

	"github.com/chon26909/apartment-api/services"
	"github.com/labstack/echo/v4"
)

type IRoomHandler interface {
	GetRooms(ctx echo.Context) error
}

type roomHandler struct {
	roomService services.IRoomService
}

func NewRoomHandler(roomService services.IRoomService) IRoomHandler {
	return &roomHandler{roomService: roomService}
}

func (h *roomHandler) GetRooms(ctx echo.Context) error {

	response, err := h.roomService.GetAllRoom()
	if err != nil {
		return err
	}

	fmt.Println("response ", response)

	return nil
}
