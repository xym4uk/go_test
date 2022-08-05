package main

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	ID   uint `gorm:"primarykey"`
	Name string
}

type Balance struct {
	ID     uint `gorm:"primarykey"`
	UserId uint
	User   User
	Amount int
}

var db = getDB()

func main() {
	var user1, user2 User
	if err := db.Where(User{ID: 2}).First(&user1).Error; err != nil {
		fmt.Println("user not found")
		return
	}
	if err := db.Where(User{ID: 1}).First(&user2).Error; err != nil {
		fmt.Println("user not found")
		return
	}
	//amount := 1000
	//changeAmount(user1.ID, -amount)
	//transfer(user1.ID, user2.ID, amount)
	getAmount(1)
}

func changeAmount(userId uint, amount int) {
	var balance Balance

	res := db.Where(Balance{UserId: userId}).Attrs(Balance{Amount: amount}).FirstOrCreate(&balance)
	if res.RowsAffected == 0 {
		db.Model(&balance).Update("amount", balance.Amount+amount)
	}
	fmt.Println("created or updated")
}

func transfer(userIdFrom uint, userIdTo uint, amount int) {
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
		if err := tx.Model(&balanceFrom).Update("amount", balanceFrom.Amount-amount).Error; err != nil {
			return err
		}

		if err := tx.Model(&balanceTo).Update("amount", balanceTo.Amount+amount).Error; err != nil {
			return err
		}
		return nil
	})
}

func getAmount(userId uint) {
	var balance Balance
	db.Where(Balance{UserId: userId}).First(&balance)
	amount := float64(balance.Amount) / 100.0
	fmt.Printf("На счету %.2f рублей", amount)
}
