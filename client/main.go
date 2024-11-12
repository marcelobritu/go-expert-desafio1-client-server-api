package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Exchange struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	log.Printf("statusCode %d\n", res.StatusCode)
	if res.StatusCode == http.StatusOK {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		var exchange Exchange
		err = json.Unmarshal(body, &exchange)
		if err != nil {
			panic(err)
		}
		printJson(&exchange)
		saveInFile(&exchange)
	}
}

func saveInFile(exchange *Exchange) {
	f, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("Dólar: %v", exchange.Bid))
	if err != nil {
		panic(err)
	}
}

func printJson(exchange *Exchange) {
	err := json.NewEncoder(os.Stdout).Encode(exchange)
	if err != nil {
		panic(err)
	}
}
