package utils

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

type MessageResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Message(status bool, message string) MessageResponse {
	return MessageResponse{
		Status:  status,
		Message: message,
	}
}

func Respond(w http.ResponseWriter, response MessageResponse) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response.Data)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

type CurrencyResponse struct {
	Rates map[string]float64 `json:"rates"`
}

func GetCurrency(currency string, currencyChanel chan<- float64) {
	defer close(currencyChanel)

	if currency != "" {
		client := http.Client{Timeout: 10 * time.Second}
		resp, err := client.Get("https://www.cbr-xml-daily.ru/latest.js")
		if err != nil {
			log.Err(err).Msg("no response from currency host")
			currencyChanel <- 1
			return
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
