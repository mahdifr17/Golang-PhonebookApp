package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mahdifr17/phonebook-v2/config"
	"github.com/mahdifr17/phonebook-v2/orm"
)

var (
	recordsOrm = orm.RecordsORM{DB: config.GetDbInstance()}
)

// GetAllRecord handle request from endpoint "/" method GET
// Used to get all records
func GetAllRecord(res http.ResponseWriter, req *http.Request) {
	// Write phonebook as JSON to response

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(recordsOrm.GetAllRecords())
}
