package server

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer(port string) {
	log.Println("Starting server in port", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
