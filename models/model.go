package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vamshireddy02/go-postgres/config"
)


type Stock struct {
	StockID int64 `json:"stockid"`
	Name string `json:"name"`
	Price int64 `json:"price"`
	Company string `json:"company"`
}


func InsertStock(stock Stock) int64 {
	db := config.CreateConnection()
	defer db.Close()

	query := `INSERT INTO stocks (name, price, company) VALUES ($1, $2, $3) RETURNING stockid`

	var id int64

	err := db.QueryRow(query, stock.Name, stock.Price, stock.Company).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a row %v",id)

	return id
}

func GetAllStocks() ([]Stock, error) {
	db := config.CreateConnection()

	defer db.Close()

	var stocks []Stock

	query := `SELECT * FROM stocks`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var stock Stock
		err := rows.Scan(&stock.Name, &stock.Price, &stock.Company)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		stocks = append(stocks, stock)

	}

	return stocks, err
}

func GetStock(id int64) (Stock, error) {

	db := config.CreateConnection()

	defer db.Close()

	var stock Stock

	query := `SELECT * FROM stocks WHERE stockid=$1`

	row := db.QueryRow(query, id)
	err := row.Scan(&stock.Name, &stock.Price, &stock.Company)
	
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return stock, nil
	case nil:
		return stock, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}
	return stock, err
}

func DeleteStock(id int64) int64 {
	db := config.CreateConnection()

	defer db.Close()

	query := `DELETE FROM stocks WHERE stockid=$1`


	res, err := db.Exec(query, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}