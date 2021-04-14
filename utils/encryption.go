package utils

import (
    "golang.org/x/crypto/bcrypt"
    log "github.com/sirupsen/logrus"
)

func HashPassword(password *string) {
    bytePass := []byte(*password)
    hashedPassword, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
    *password = string(hashedPassword)
}

func ComparePassword(hashedPassword, password string){
    hash := []byte(hashedPassword)
    pass := []byte(password)
    err := bcrypt.CompareHashAndPassword(hash, pass)
    if err != nil {
        log.Errorln("Couldn't Compare Hash and Password")
    }
}
