package controllers

import (
	"github.com/gorilla/mux"
	"github.com/xym4uk/testAvito/models"
	u "github.com/xym4uk/testAvito/utils"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type TransactionControllerImpl struct {
	db *gorm.DB
}

func NewTransactionController(db *gorm.DB) TransactionControllerImpl {
	return TransactionControllerImpl{db: db}
}

func (tc TransactionControllerImpl) GetTransactions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["userId"])
	if err != nil {
		u.Respond(w, u.Message(false, "Некорректный id"))
		return
	}

	data := models.GetTransactions(ctx, uint(id), tc.db)

	resp := u.Message(true, "success")
	resp.Data = data
	u.Respond(w, resp)
}
