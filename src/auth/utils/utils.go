package utils

import "encoding/json"

func ToJSON(model *interface{}) ([]byte, error) {
	data, err := json.Marshal(&model)
	return data, err
}
