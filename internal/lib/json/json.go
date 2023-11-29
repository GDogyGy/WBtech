package json

import (
	"WbTech0/internal/model"
	"encoding/json"
)

func ParseToModel(b []byte) (model.Order, error) {
	var date model.Order
	if err := json.Unmarshal(b, &date); err != nil {
		return model.Order{}, err
	}

	return date, nil
}
