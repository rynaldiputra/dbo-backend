package order

import (
	"dbo-be/entities"
)

type OrderFormatter struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

func FormatOrder(order entities.Order) OrderFormatter {
	formatter := OrderFormatter{
		ID:       order.ID,
		UserID:   order.UserID,
		Name:     order.Name,
		Type:     order.Type,
		Quantity: order.Quantity,
		Price:    order.Price,
	}

	return formatter
}
