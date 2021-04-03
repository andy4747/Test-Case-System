package handlers

import (
	"encoding/json"
	"github.com/angeldhakal/testcase-ms/models"
	"github.com/angeldhakal/testcase-ms/utils"
	"log"
	"net/http"
)

func createTestCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var testCase models.TestCaseModel
	err := json.NewDecoder(r.Body).Decode(&testCase)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	log.Printf("%+v",testCase)
	conn := models.Connect()
	createTestCaseQuery := `INSERT INTO test_cases (title, date, tested_by, functionality, summary, description, data, url, expected_result, actual_result, environment, device) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`
	_, err = conn.Exec(createTestCaseQuery, testCase.Title, testCase.Date, testCase.TestedBy, testCase.Functionality,testCase.Summary,testCase.Description,testCase.Data,testCase.URL,testCase.ExpectedResult,testCase.ActualResult,testCase.Environment,testCase.Device)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)

	// debugging code
	var jsonStr []byte
	jsonStr, err = utils.ParseToJson(&testCase)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	w.Write(jsonStr)
}
