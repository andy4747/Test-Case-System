package cases_handler

import (
	"github.com/angeldhakal/testcase-ms/models"
	"github.com/gorilla/mux"
	log	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func deleteTestCase(w http.ResponseWriter, r *http.Request) {
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

	deleteSqlQuery := `DELETE FROM test_cases WHERE id=$1`
	_, err = conn.Exec(deleteSqlQuery, id)
	if err != nil {
		log.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
