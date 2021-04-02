package utils

import "encoding/json"

func ParseToJson(v interface{}) ([]byte, error) {
	jsonStr, err := json.MarshalIndent(&v,"","  ")
	if err != nil {
		return nil, err
	}
	return jsonStr, nil
}
