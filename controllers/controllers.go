package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vamshireddy02/go-postgres/models"
)


type response struct {
	ID int64 `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}


func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		fmt.Println("Unable to decode the request body.")
	}

	insertID := models.InsertStock(stock)

	res := response{
		ID : insertID,
		Message: "stock is created",
	}

	json.NewEncoder(w).Encode(res)

}

func GetAllStock(w http.ResponseWriter, r *http.Request) {

	stocks, err := models.GetAllStocks()
	if err != nil {
		log.Fatalf("Unable to get all stock. %v", err)
	}

	json.NewEncoder(w).Encode(stocks)
}

func GetStock(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	stock, err := models.GetStock(int64(id))

	if err != nil {
		log.Fatalf("Unable to get stock. %v", err)
	}

	json.NewEncoder(w).Encode(stock)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRow := models.DeleteStock(int64(id))
	msg := fmt.Sprintf("Stock updated successfully. Total rows/record affected %v", deletedRow)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}