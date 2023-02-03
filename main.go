package main

import (
	"github.com/enfipy/locker"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/xym4uk/testAvito/controllers"
	"github.com/xym4uk/testAvito/utils"
	"net/http"
)

func main() {
	db := utils.Init()
	lock := locker.Initialize()
	log.Log().Msg("Starting http server...")

	router := mux.NewRouter()

	balanceController := controllers.NewBalanceController(db, lock)
	transactionController := controllers.NewTransactionController(db)
	// получить баланс пользователя
	router.HandleFunc("/api/balance/{userId}", balanceController.GetBalance).Methods("GET")
	// перевод средств между пользователями
	router.HandleFunc("/api/balance/transfer", balanceController.Transfer).Methods("POST")
	// начисление\списание средств с баланса
	router.HandleFunc("/api/balance/change-amount", balanceController.ChangeAmount).Methods("POST")
	// получение транзакций пользователя
	router.HandleFunc("/api/transactions/{userId}", transactionController.GetTransactions).Methods("GET")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Err(err).Msg("")
		return
	}
}
