package models

import (
	"github.com/xym4uk/testAvito/utils"
	"time"
)

type Transaction struct {
	ID        uint `gorm:"primarykey"`
	Amount    int
	Comment   string
	UserID    uint
	CreatedAt time.Time
}

func GetTransactions(userId uint) *[]Transaction {
	db, _ := utils.GetDB()
	var transactions []Transaction
	db.Where(Transaction{UserID: userId}).Find(&transactions)

	return &transactions
}
