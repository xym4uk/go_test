package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/xym4uk/testAvito/controllers"
	"github.com/xym4uk/testAvito/models"
	"net/http"
)

func main() {
	var _, dbErr = models.GetDB()
	if dbErr != nil {
		log.Fatal().Err(dbErr).Msg("")
		return
	}
	log.Log().Msg("Starting http server...")

	router := mux.NewRouter()

	// получить баланс пользователя
	router.HandleFunc("/api/balance/{userId}", controllers.GetBalance).Methods("GET")
	// перевод средств между пользователями
	router.HandleFunc("/api/balance/transfer", controllers.Transfer).Methods("POST")
	// начисление\списание средств с баланса
	router.HandleFunc("/api/balance/change-amount", controllers.ChangeAmount).Methods("POST")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Err(err).Msg("")
		return
	}
}
