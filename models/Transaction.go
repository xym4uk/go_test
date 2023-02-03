package models

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID        uint `gorm:"primarykey"`
	Amount    int
	Comment   string
	UserID    uint
	CreatedAt time.Time
}

func GetTransactions(ctx context.Context, userId uint, db *gorm.DB) *[]Transaction {
	var transactions []Transaction
	db.WithContext(ctx).Where(Transaction{UserID: userId}).Find(&transactions)

	return &transactions
}
