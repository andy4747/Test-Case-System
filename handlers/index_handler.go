package handlers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type apiPaths struct {
	Path string `json:"path"`
	Functionality string `json:"functionality"`
}

func allPaths() []apiPaths {
	paths := []apiPaths{
		{"/", "returns all endpoints of this service."},
		{"/case/create", "POST endpoint that create a test case."},
		{"/cases", "gets all the testcases."},
		{"/case/delete/{id}", "POST endpoint that create a test case."},
	}
	return paths
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var paths []apiPaths = allPaths()

	pathsJson, err := json.MarshalIndent(paths, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorln(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(pathsJson)
}
