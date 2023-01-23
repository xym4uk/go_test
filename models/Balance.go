package models

import (
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

type Balance struct {
	ID     uint `gorm:"primarykey"`
	UserId uint
	User   User
	Amount int
}

func GetAmount(userId uint) *Balance {
	var balance Balance
	db, _ := GetDB()
	db.Where(Balance{UserId: userId}).First(&balance)
	amount := float64(balance.Amount) / 100.0
	fmt.Printf("На счету %.2f рублей", amount)

	return &balance
}

func Transfer(userIdFrom uint, userIdTo uint, amount int) {
	db, _ := GetDB()
	var balanceFrom, balanceTo Balance
	if amount <= 0 {
		return
	}
	db.Where(Balance{UserId: userIdFrom}).FirstOrCreate(&balanceFrom)
	if balanceFrom.Amount < amount {
		fmt.Println("Недостаточно бабосов")
		return
	}

	db.Where(Balance{UserId: userIdTo}).Attrs(Balance{Amount: 0}).FirstOrCreate(&balanceTo)
	db.Transaction(func(tx *gorm.DB) error {
		var transaction Transaction
		if err := tx.Model(&balanceFrom).Update("amount", balanceFrom.Amount-amount).Error; err != nil {
			return err
		}

		if err := tx.Model(&balanceTo).Update("amount", balanceTo.Amount+amount).Error; err != nil {
			return err
		}

		transaction.Amount = amount
		transaction.Comment = "Перевод пользователю. ID: " + strconv.FormatUint(uint64(balanceTo.UserId), 10)
		tx.Create(&transaction)

		return nil
	})
}

func ChangeAmount(userId uint, amount int) {
	db, _ := GetDB()
	var balance Balance

	res := db.Where(Balance{UserId: userId}).Attrs(Balance{Amount: amount}).FirstOrCreate(&balance)
	if res.RowsAffected == 0 {
		db.Model(&balance).Update("amount", balance.Amount+amount)
	}
	fmt.Println("created or updated")
}
