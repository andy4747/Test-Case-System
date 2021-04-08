package cases_handler

import (
	"encoding/json"
	"github.com/angeldhakal/testcase-ms/models"
	log "github.com/sirupsen/logrus"
	"net/http"
)


func getTestCases(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	conn := models.Connect()
	getAllQuery := `SELECT * FROM test_cases`
	log.Infoln(getAllQuery)
	queryRows, err := conn.Query(getAllQuery)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorln(err)
		return
	}
	var cases []models.TestCaseModel
	for queryRows.Next() {
		var testCase models.TestCaseModel
		err := queryRows.Scan(&testCase.ID,&testCase.Title, &testCase.Date, &testCase.TestedBy, &testCase.Functionality,&testCase.Summary,&testCase.Description,&testCase.Data,&testCase.URL,&testCase.ExpectedResult,&testCase.ActualResult,&testCase.Environment,&testCase.Device)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Errorln(err)
			return
		}
		cases = append(cases,testCase)
	}
	var jsonCases []byte
	jsonCases, err = json.MarshalIndent(&cases,"","  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorln(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonCases)
}
