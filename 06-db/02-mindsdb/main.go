package main

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "mindsdb@tcp(127.0.0.1:47335)/mindsdb")
	check(err)
	if err != nil {
		return
	}
	defer db.Close()
	err = db.Ping()
	check(err)
	if err != nil {
		return
	}
	predict()
}

type Prediction struct {
	PredictedValue float64 `json:"predicted"`
	Confidence     float64 `json:"confidence"`
}

func predict() error {
	log.Printf("Predicting the DELIVERED_ST.\n")
	pred := Prediction{}
	err := db.QueryRow(`SELECT DELIVERED_ST, DELIVERED_ST_confidence as CONFIDENCE FROM sprints_delivered_st_predictor WHERE when_data='{
             "TEAM_NAME": "100-XECM"
	}'`).Scan(&pred.PredictedValue, &pred.Confidence)
	check(err)
	if err != nil {
		return err
	}
	result, _ := json.MarshalIndent(&pred, "", " ")
	log.Printf("Result: %s\n", string(result))
	return nil
}

func check(err error) {
	if err != nil {
		log.Printf("Error: %s", err.Error())
	}
}
