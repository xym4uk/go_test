package controllers

import (
	"encoding/json"
	"github.com/enfipy/locker"
	"github.com/gorilla/mux"
	"github.com/xym4uk/testAvito/controllers/requests"
	"github.com/xym4uk/testAvito/models"
	u "github.com/xym4uk/testAvito/utils"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type BalanceControllerImpl struct {
	db *gorm.DB
	mu *locker.Locker
}

func NewBalanceController(db *gorm.DB, mu *locker.Locker) BalanceControllerImpl {
	return BalanceControllerImpl{
		db: db,
		mu: mu,
	}
}

func (bc BalanceControllerImpl) GetBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var currencyChanel = make(chan float64, 1)
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["userId"])
	if err != nil {
		u.Respond(w, u.Message(false, "Некорректный id"))
		return
	}

	go u.GetCurrency(r.FormValue("currency"), currencyChanel)

	data := models.GetAmount(ctx, uint(id), bc.db)
	cur := <-currencyChanel
	data.Amount = int(float64(data.Amount) * cur)
	resp := u.Message(true, "success")
	resp.Data = data
	u.Respond(w, resp)
}

func (bc BalanceControllerImpl) Transfer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	decoder := json.NewDecoder(r.Body)
	var tr requests.TransferRequest
	err := decoder.Decode(&tr)
	if err != nil {
		u.Respond(w, u.Message(false, "invalid request"))
		return
	}
	bc.mu.Lock(strconv.FormatUint(uint64(tr.From), 10))
	models.Transfer(ctx, tr.From, tr.To, tr.Amount, bc.db)
	bc.mu.Unlock(strconv.FormatUint(uint64(tr.From), 10))

	u.Respond(w, u.Message(true, "success"))
}

func (bc BalanceControllerImpl) ChangeAmount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	decoder := json.NewDecoder(r.Body)
	var c requests.ChangeAmountRequest
	err := decoder.Decode(&c)
	if err != nil {
		u.Respond(w, u.Message(false, "invalid request"))
		return
	}
	lock := locker.Initialize()
	lock.Lock(strconv.FormatUint(uint64(c.UserID), 10))
	models.ChangeAmount(ctx, c.UserID, c.Amount, bc.db)
	lock.Unlock(strconv.FormatUint(uint64(c.UserID), 10))

	u.Respond(w, u.Message(true, "success"))
}
