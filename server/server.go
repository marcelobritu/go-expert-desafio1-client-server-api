package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func ApiGetDollarValue() string {
	ctxApi, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctxApi, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}

func StartServer(port string) {
	log.Println(ApiGetDollarValue())
	log.Println("Starting server in port", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
