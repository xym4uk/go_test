package utils

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

type CurrencyResponse struct {
	Rates map[string]float64 `json:"rates"`
}

func GetCurrency(currency string, currencyChanel chan<- float64) {
	defer close(currencyChanel)

	if currency != "" {
		resp, err := http.Get("https://www.cbr-xml-daily.ru/latest.js")
		if err != nil {
			log.Err(err)
		}

		var result CurrencyResponse
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			currencyChanel <- 1
			return
		}

		currencyRate := result.Rates[currency]
		if currencyRate != 0 {
			currencyChanel <- currencyRate
			return
		}
	}

	currencyChanel <- 1
}
