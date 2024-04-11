package models

import (
	"time"
)

// database
type Room struct {
	Id             int       `gorm:"column:id"`
	Building       int       `gorm:"column:building"`
	Floor          int       `gorm:"column:floor"`
	Area           string    `gorm:"column:area"`
	NumberOfPeople int       `gorm:"column:number_of_people"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	CreatedBy      string    `gorm:"column:created_by"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
	UpdatedBy      string    `gorm:"column:updated_by"`
}

// get all rooms
type GetRoomsResponse struct {
	Response
	Data []Room `json:"data"`
}
