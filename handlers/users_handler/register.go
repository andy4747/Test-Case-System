package users_handler

import (
	"encoding/json"
	"github.com/angeldhakal/testcase-ms/models"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var id uint64 = 0
	var user models.Users
	user.ID = id
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Errorln(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

