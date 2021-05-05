package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/angeldhakal/testcase-ms/models"
	"github.com/angeldhakal/testcase-ms/repository"
	"github.com/angeldhakal/testcase-ms/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type CaseHandler interface {
	GetCase(http.HandlerFunc)
	CreateCase(http.HandlerFunc)
	UpdateCase(http.HandlerFunc)
	DeleteCase(http.HandlerFunc)
}

type caseHandler struct {
	repo repository.CaseRepository
}

func NewCaseHandler() caseHandler {
	return caseHandler{
		repo: repository.NewCaseRepository(),
	}
}

func (h *caseHandler) GetCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Errorln("Invalid id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	token := w.Header().Get("token")
	if token == "" {
		log.Errorln("Enter a auth token in the header")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var tokenClaim *jwt.Token
	tokenClaim, err = util.ValidateToken(token)
	if err != nil {
		log.Errorln("Invalid Token")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	claims, ok := tokenClaim.Claims.(jwt.MapClaims)
	if !ok {
		log.Errorln("Couldn't successfully retrieve the token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user := claims["user"]
	var testCase models.TestCaseModel
	testCase, err = h.repo.GetCase(id, user.(string))
	if err != nil {
		log.Errorln("No data found in the database")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var caseJson []byte
	caseJson, err = json.MarshalIndent(testCase, "", "  ")
	if err != nil {
		log.Errorln("Couldn't marshal json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(caseJson)
}

func (h *caseHandler) CreateCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var testCase models.TestCaseModel
	err := json.NewDecoder(r.Body).Decode(&testCase)
	if err != nil {
		log.Errorln("Invalid Data Entered")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = h.repo.AddCase(testCase)
	if err != nil {
		log.Errorln(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *caseHandler) UpdateCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var testCase models.TestCaseModel
	err := json.NewDecoder(r.Body).Decode(&testCase)
	if err != nil {
		log.Errorln(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = h.repo.UpdateCase(testCase)
	if err != nil {
		log.Errorln(err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *caseHandler) DeleteCase(w http.ResponseWriter, r *http.Request) {

}
