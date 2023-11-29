package filterOrder

import "WbTech0/internal/model"

func GetOrderById(data []model.Order, id string) model.Order {
	for _, v := range data {
		if v.OrderUid == id {
			return v
		}
	}
	return model.Order{}
}
