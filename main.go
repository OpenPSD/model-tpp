package main

import (
	"log"
	"net/http"

	"github.com/openpsd/model-tpp/api"
)

func main() {
	log.Println("start OpenPSD DeliveryThinking TPP reference implementation server")
	s := api.TppHTTPServer{BankID: "mobinauten", PIN: "02011"}

	TppAPI := api.NewMockedTppHTTPServer(s)
	http.ListenAndServe(":8080", TppAPI)
}
