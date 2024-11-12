package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

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

var ErrTimeout = errors.New("request timeout exceeded")

func ExchangeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cotacao" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	exchange, err := GetExchangeRate()
	if err != nil {
		log.Printf("Error fetching exchange rate: %v", err)
		if errors.Is(err, ErrTimeout) {
			w.WriteHeader(http.StatusGatewayTimeout)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(exchange.Usdbrl)
}

func GetExchangeRate() (*Exchange, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return nil, ErrTimeout
		}
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var exchange Exchange
	err = json.Unmarshal(body, &exchange)
	if err != nil {
		return nil, err
	}

	return &exchange, nil
}

func main() {
	log.Println("Starting server in port 8080")
	http.HandleFunc("/cotacao", ExchangeHandler)
	http.ListenAndServe(":8080", nil)
}
