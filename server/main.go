package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type CotacaoResponse struct {
	Bid string `json:"bid"`
}

type Exchange struct {
	Usdbrl struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func ExchangeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cotacao" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	exchange, err := GetExchangeRate("USD-BRL")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	cotacaoResponse := CotacaoResponse{Bid: exchange.Usdbrl.Bid}
	json.NewEncoder(w).Encode(cotacaoResponse)
}

func GetExchangeRate(currency string) (*Exchange, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/"+currency, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var c Exchange
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func main() {
	log.Println("Starting server in port 8080")
	http.HandleFunc("/cotacao", ExchangeHandler)
	http.ListenAndServe(":8080", nil)
}
