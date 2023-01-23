package controllers

import (
	"github.com/gorilla/mux"
	"github.com/xym4uk/testAvito/models"
	u "github.com/xym4uk/testAvito/utils"
	"net/http"
	"strconv"
)

var GetTransactions = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["userId"])
	if err != nil {
		u.Respond(w, u.Message(false, "Некорректный id"))
	}

	data := models.GetTransactions(uint(id))

	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
