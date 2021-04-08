package cases_handler

import (
	"encoding/json"
	"github.com/angeldhakal/testcase-ms/models"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func  getTestCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"],10,64)
	if err != nil {
		log.Errorln("{id} should be a positive integer.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//establishing database connection
	conn := models.Connect()

	// check if record exists in the database using id
	checkRecordQuery := `SELECT EXISTS(SELECT id FROM test_cases WHERE id=$1)`
	var exists bool

	err = conn.QueryRow(checkRecordQuery, id).Scan(&exists)
	if err != nil {
		log.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if exists != true {
		log.Errorln("{id} doesn't exists")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//selecting record from database
	var testCase models.TestCaseModel
	getCaseQuery := `SELECT * FROM test_cases WHERE id=$1`
	err = conn.QueryRow(getCaseQuery, id).Scan(&testCase.ID,&testCase.Title, &testCase.Date, &testCase.TestedBy, &testCase.Functionality,&testCase.Summary,&testCase.Description,&testCase.Data,&testCase.URL,&testCase.ExpectedResult,&testCase.ActualResult,&testCase.Environment,&testCase.Device)
	if err != nil {
		log.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var jsonCase []byte
	jsonCase, err = json.MarshalIndent(&testCase,"","  ")
	if err != nil {
		log.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonCase)
}