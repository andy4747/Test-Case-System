package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/angeldhakal/testcase-ms/models"
	"github.com/angeldhakal/testcase-ms/repository"
	"github.com/angeldhakal/testcase-ms/utils"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// TODOS
/*
TODO => add current time in AddUser Handler for CurrentDate Field
*/

type UserHandler interface {
	AddUser(http.HandlerFunc)
	GetUser(http.HandlerFunc)
	GetAllUser(http.HandlerFunc)
	UpdateUser(http.HandlerFunc)
	DeleteUser(http.HandlerFunc)
	SignInUser(http.HandlerFunc)
}

type userHandler struct {
	repo repository.UserRepository
}

func NewUserHandler() userHandler {
	return userHandler{
		repo: repository.NewUserRepository(),
	}
}

func (h *userHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Errorln("Invalid Request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	utils.HashPassword(&user.Password)
	user, err = h.repo.AddUser(user)
	if err != nil {
		log.Errorln("Couldn't Register user due to some server issues")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user.Password = ""
	w.WriteHeader(http.StatusCreated)
}

func (h *userHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.Users
	// parsing id parameter from url endpoint to integer value
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Warnln("Enter valid user id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err = h.repo.GetUser(id)
	if err != nil {
		log.Warnln("Couldn't get user due to some server issue")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var userJson []byte
	userJson, err = json.MarshalIndent(&user, "", "  ")
	if err != nil {
		log.Errorln("Couldn't encode json due to some server issues")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(userJson)
}

func (h *userHandler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.Users
	users, err := h.repo.GetAllUsers()
	if err != nil {
		log.Errorln("Couldn't get users due to some server issues")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var allUsers []byte
	allUsers, err = json.MarshalIndent(&users, "", "  ")
	if err != nil {
		log.Errorln("Could't encode the json due to some server issue")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(allUsers)
}

func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Errorln("Provide valid information")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err = h.repo.UpdateUser(user)
	if err != nil {
		log.Errorln("Couldn't update due to some server issues")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Errorln("Enter valid user id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var user models.Users
	user.ID = uint(id)
	user, err = h.repo.DeleteUser(user)
	if err != nil {
		log.Errorln("Couldn't delete user due to some server error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *userHandler) SignInUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user utils.Credentials
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Errorln("Enter valid login credentials")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dbUser, err := h.repo.GetUserByEmail(user.Email)
	if err != nil {
		log.Errorln("Couldn't get user from database due to some server issues")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isCorrectPassword := utils.ComparePassword(dbUser.Password, user.Password); isCorrectPassword {
		token, err := utils.GenerateToken(dbUser.ID)
        if err != nil {
            log.Errorln("Token generation failed")
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
		cookie := http.Cookie{
			Name:     "token",
			Value:    token,
			HttpOnly: true,
			Secure:   true,
		}
		log.Infoln(token)
		http.SetCookie(w, &cookie)
		w.Header().Add("token", token)
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	return
}
