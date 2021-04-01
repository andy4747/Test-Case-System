package handlers

import (
	"encoding/json"
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
	}
	return paths
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var paths []apiPaths = allPaths()

	pathsJson, err := json.MarshalIndent(paths, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(pathsJson)
}
