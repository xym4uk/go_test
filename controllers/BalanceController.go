package controllers

import (
	"encoding/json"
	"github.com/enfipy/locker"
	"github.com/gorilla/mux"
	"github.com/xym4uk/testAvito/controllers/requests"
	"github.com/xym4uk/testAvito/models"
	u "github.com/xym4uk/testAvito/utils"
	"net/http"
	"strconv"
)

var GetBalance = func(w http.ResponseWriter, r *http.Request) {
	var currencyChanel = make(chan float64, 1)
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["userId"])
	if err != nil {
		u.Respond(w, u.Message(false, "Некорректный id"))
	}

	go u.GetCurrency(r.FormValue("currency"), currencyChanel)

	cur := <-currencyChanel

	println(cur)

	data := models.GetAmount(uint(id))
	data.Amount = int(float64(data.Amount) * cur)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var Transfer = func(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var tr requests.TransferRequest
	err := decoder.Decode(&tr)
	if err != nil {
		u.Respond(w, u.Message(false, "invalid request"))
	}
	lock := locker.Initialize()
	lock.Lock(strconv.FormatUint(uint64(tr.From), 10))
	models.Transfer(tr.From, tr.To, tr.Amount)
	lock.Unlock(strconv.FormatUint(uint64(tr.From), 10))

	u.Respond(w, u.Message(true, "success"))
}

var ChangeAmount = func(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var c requests.ChangeAmountRequest
	err := decoder.Decode(&c)
	if err != nil {
		u.Respond(w, u.Message(false, "invalid request"))
	}
	lock := locker.Initialize()
	lock.Lock(strconv.FormatUint(uint64(c.UserID), 10))
	models.ChangeAmount(c.UserID, c.Amount)
	lock.Unlock(strconv.FormatUint(uint64(c.UserID), 10))

	u.Respond(w, u.Message(true, "success"))
}
