package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/xym4uk/testAvito/controllers/requests"
	"github.com/xym4uk/testAvito/models"
	u "github.com/xym4uk/testAvito/utils"
	"net/http"
	"strconv"
)

var GetBalance = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["userId"])
	if err != nil {
		u.Respond(w, u.Message(false, "Некорректный id"))
	}

	data := models.GetAmount(uint(id))

	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var Transfer = func(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t requests.TransferRequest
	err := decoder.Decode(&t)
	if err != nil {
		u.Respond(w, u.Message(false, "invalid request"))
	}
	models.Transfer(t.From, t.To, t.Amount)

	u.Respond(w, u.Message(true, "success"))
}

var ChangeAmount = func(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var c requests.ChangeAmountRequest
	err := decoder.Decode(&c)
	if err != nil {
		u.Respond(w, u.Message(false, "invalid request"))
	}
	models.ChangeAmount(c.UserID, c.Amount)

	u.Respond(w, u.Message(true, "success"))
}
