package models

import (
	"fmt"
	"testing"
)

func TestConnection(t *testing.T) {
	conn := Connect()
	defer conn.Close()
	err := conn.Ping()
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println("Pong")
}
