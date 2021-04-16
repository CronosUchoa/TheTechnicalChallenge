package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//estrutura para receber o json

type blockCoin struct {
	Txid          string `json:"txid"`
	Vout          int    `json:"vout"`
	Height        int    `json:"height"`
	Value         string `json:"value"`
	Confirmations int    `json:"confirmations"`
}

// func getUrl(ur[]) {

// 	fmt.Sprintf("https://blockbook-bitcoin.tronwallet.me/api/v2/utxo/%s", ur)

// }

func getCoins(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	r.ParseMultipartForm(0)

	address := r.FormValue("address")

	params := url.Values{}
	params.Add("address", address)
	str := strings.Join(params["address"], "")
	// boddy := strings.NewReader(params.Encode())
	p := fmt.Sprintf("https://blockbook-bitcoin.tronwallet.me/api/v2/utxo/%s", str)

	resp, err := http.Get(p)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Fprintf(w, "%s", params["address"])
	//response seria a resposta do servidor a requisição
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(body)
	}

	var content []blockCoin

	err = json.Unmarshal(body, &content)
	if err != nil {
		log.Fatal(err)
	}

	// simpleObjectJson, err := json.Marshal(content)
	// if err != nil {
	// 	fmt.Println("Erro!")
	// }
	json.NewEncoder(w).Encode(content)

}

func main() {

	fmt.Println("Calling API port 8000...")

	r := mux.NewRouter().StrictSlash(true)

	headers := handlers.AllowedHeaders([]string{"X-Request", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	r.HandleFunc("/balance/", getCoins).Methods("GET")
	

	// http.ListenAndServe(":8000", nil)
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins)(r)))

}
