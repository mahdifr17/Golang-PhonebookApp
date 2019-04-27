package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mahdifr17/phonebook-v2/exception"
	"github.com/mahdifr17/phonebook-v2/structs"

	"github.com/gorilla/mux"

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
	json.NewEncoder(res).Encode(structs.BaseResponse{Code: 200, Msg: "OK", Result: recordsOrm.GetAllRecords()})
}

// GetRecord handle request from endpoint "/{id}" method GET
// Used to retrive specific record
func GetRecord(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req) // Get request param
	if i, errParse := strconv.Atoi(params["id"]); errParse == nil {
		if rec, errExist := recordsOrm.GetRecord(i); errExist == nil {
			json.NewEncoder(res).Encode(structs.BaseResponse{Code: 200, Msg: "OK", Result: rec})
		} else {
			// Record not found
			json.NewEncoder(res).Encode(exception.SerializeError(errExist))
		}
	} else {
		// Id not instanceOf int
		json.NewEncoder(res).Encode(exception.SerializeError(errParse))
	}
}

// AddRecord handle request from endpoint "/" method POST
// Used to add single record
func AddRecord(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var newRecord structs.Record
	// Decode request body, store to record
	json.NewDecoder(req.Body).Decode(&newRecord)

	if result, err := recordsOrm.AddRecord(newRecord); err == nil {
		json.NewEncoder(res).Encode(structs.BaseResponse{Code: 200, Msg: "OK", Result: result})
	} else {
		// Error insert
		json.NewEncoder(res).Encode(exception.SerializeError(err))
	}
}

// UpdateRecord handle request from endpoint "/{id}" method PUT
// Used to update record
func UpdateRecord(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req) // Get request param
	if i, errParse := strconv.Atoi(params["id"]); errParse == nil {
		if rec, errExist := recordsOrm.GetRecord(i); errExist == nil {
			var inputRecord structs.Record
			// Decode request body, store to inputRecord
			json.NewDecoder(req.Body).Decode(&inputRecord)

			// Update changed value
			if result, errUpdate := recordsOrm.UpdateRecord(rec, inputRecord); errUpdate == nil {
				json.NewEncoder(res).Encode(structs.BaseResponse{Code: 200, Msg: "OK", Result: result})
			} else {
				// Error update
				json.NewEncoder(res).Encode(exception.SerializeError(errUpdate))
			}
		} else {
			// Record not found
			json.NewEncoder(res).Encode(exception.SerializeError(errExist))
		}
	} else {
		// Id not instanceOf int
		json.NewEncoder(res).Encode(exception.SerializeError(errParse))
	}
}

// DeleteRecord handle request from endpoint "/{id}" method DELETE
// Used to delete record
func DeleteRecord(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req) // Get request param
	if i, errParse := strconv.Atoi(params["id"]); errParse == nil {
		if rec, errExist := recordsOrm.GetRecord(i); errExist == nil {
			if result, errDelete := recordsOrm.DeleteRecord(rec); errDelete == nil {
				json.NewEncoder(res).Encode(structs.BaseResponse{Code: 200, Msg: "OK", Result: result})
			} else {
				// Error update
				json.NewEncoder(res).Encode(exception.SerializeError(errDelete))
			}
		} else {
			// Record not found
			json.NewEncoder(res).Encode(exception.SerializeError(errExist))
		}
	} else {
		// Id not instanceOf int
		json.NewEncoder(res).Encode(errParse.Error())
	}
}
