package models

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Balance struct {
	ID     uint `gorm:"primarykey"`
	UserID uint `gorm:"column:user_id;type:uint REFERENCES users(id);"`
	User   User `gorm:"foreignKey:UserID;references:ID"`
	Amount int
}

func GetAmount(ctx context.Context, userId uint, db *gorm.DB) *Balance {
	var balance Balance
	db.WithContext(ctx).Preload("User").Where(Balance{UserID: userId}).First(&balance)

	return &balance
}

func Transfer(ctx context.Context, userIdFrom uint, userIdTo uint, amount int, db *gorm.DB) {
	var balanceFrom, balanceTo Balance
	if amount <= 0 {
		return
	}
	db.WithContext(ctx).Where(Balance{UserID: userIdFrom}).FirstOrCreate(&balanceFrom)
	if balanceFrom.Amount < amount {
		fmt.Println("Недостаточно средств")
		return
	}

	db.WithContext(ctx).Where(Balance{UserID: userIdTo}).Attrs(Balance{Amount: 0}).FirstOrCreate(&balanceTo)
	err := db.Transaction(func(tx *gorm.DB) error {
		var transaction Transaction
		if err := tx.WithContext(ctx).Model(&balanceFrom).Update("amount", balanceFrom.Amount-amount).Error; err != nil {
			return err
		}

		if err := tx.WithContext(ctx).Model(&balanceTo).Update("amount", balanceTo.Amount+amount).Error; err != nil {
			return err
		}

		transaction.Amount = amount
		transaction.Comment = "Перевод пользователю. ID: " + strconv.FormatUint(uint64(balanceTo.UserID), 10)
		transaction.UserID = balanceFrom.UserID
		transaction.CreatedAt = time.Now()
		tx.WithContext(ctx).Create(&transaction)

		return nil
	})
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

func ChangeAmount(ctx context.Context, userId uint, amount int, db *gorm.DB) {
	var balance Balance
	var transaction Transaction

	res := db.WithContext(ctx).Where(Balance{UserID: userId}).Attrs(Balance{Amount: amount}).FirstOrCreate(&balance)
	if res.RowsAffected == 0 {
		db.WithContext(ctx).Model(&balance).Update("amount", balance.Amount+amount)
	}

	transaction.Amount = amount
	transaction.Comment = "Изменение баланса: " + strconv.Itoa(amount)
	transaction.UserID = userId
	transaction.CreatedAt = time.Now()
	db.WithContext(ctx).Create(&transaction)
}
