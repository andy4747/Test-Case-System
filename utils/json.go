package utils

import (
	"encoding/json"
	"log"
)

func ParseToJson(v interface{}) ([]byte, error) {
	jsonStr, err := json.MarshalIndent(&v,"","  ")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return jsonStr, nil
}