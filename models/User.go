package models

type User struct {
	ID   uint `gorm:"primarykey"`
	Name string
}
