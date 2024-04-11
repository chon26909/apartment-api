package models

type User struct {
	Id        int    `gorm:"column:id"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Age       int    `gorm:"column:age"`
	Email     string `gorm:"column:email"`
	Hash      string `gorm:"column:hash"`
	Salt      string `gorm:"column:salt"`
}
