package utils

import (
	"encoding/json"
)

func Encode(object any) (string, error) {
	json, err := json.Marshal(object)
	if err != nil {
		return "", err
	}
	return string(json), nil
}
