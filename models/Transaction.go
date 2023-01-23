package models

type Transaction struct {
	ID      uint `gorm:"primarykey"`
	Amount  int
	Comment string
}
