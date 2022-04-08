package transformers

import (
	"time"

	"dbo.id/product-service/models"
)

type (
	Order struct {
		Id                int         `json:"id"`
		TransactionNumber string      `json:"transaction_number"`
		Media             string      `json:"media"`
		IsMember          int         `json:"is_member"`
		CustomerName      string      `json:"customer_name"`
		CustomerEmail     string      `json:"customer_email"`
		CustomerPhone     string      `json:"customer_phone"`
		CustomerAddress   string      `json:"customer_address"`
		OrderItems        []OrderItem `json:"items"`
		CreatedAt         time.Time   `json:"created_at"`
		UpdatedAt         time.Time   `json:"updated_at"`
	}

	OrderItem struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}
)

func (res *CollectionPagingTransformer) TransformOrderList(arr []models.Order, pagination *models.Pagination) {
	for _, item := range arr {
		order := Order{}
		order.Id = item.Id
		order.TransactionNumber = item.TransactionNumber
		order.Media = item.Media
		order.IsMember = item.IsMember
		order.CustomerName = item.CustomerName
		order.CustomerEmail = item.CustomerEmail
		order.CustomerPhone = item.CustomerPhone
		order.CustomerAddress = item.CustomerAddress

		if item.IsMember == 1 {
			order.CustomerName = item.Customer.Name
			order.CustomerEmail = item.Customer.Email
			order.CustomerPhone = item.Customer.Phone
			order.CustomerAddress = item.Customer.Address
		}

		for _, orderItemResult := range item.OrderItems {
			orderItem := OrderItem{}
			orderItem.Name = orderItemResult.Name
			orderItem.Price = orderItemResult.Price

			order.OrderItems = append(order.OrderItems, orderItem)
		}

		order.CreatedAt = item.CreatedAt
		order.UpdatedAt = item.UpdatedAt

		res.Data = append(res.Data, order)
	}

	res.Meta = pagination
}

func (res *Transformer) TransformOrderGetById(item models.Order) {
	order := Order{}
	order.Id = item.Id
	order.TransactionNumber = item.TransactionNumber
	order.Media = item.Media
	order.IsMember = item.IsMember
	order.CustomerName = item.CustomerName
	order.CustomerEmail = item.CustomerEmail
	order.CustomerPhone = item.CustomerPhone
	order.CustomerAddress = item.CustomerAddress

	if item.IsMember == 1 {
		order.CustomerName = item.Customer.Name
		order.CustomerEmail = item.Customer.Email
		order.CustomerPhone = item.Customer.Phone
		order.CustomerAddress = item.Customer.Address
	}

	for _, orderItemResult := range item.OrderItems {
		orderItem := OrderItem{}
		orderItem.Name = orderItemResult.Name
		orderItem.Price = orderItemResult.Price

		order.OrderItems = append(order.OrderItems, orderItem)
	}

	order.CreatedAt = item.CreatedAt
	order.UpdatedAt = item.UpdatedAt

	res.Data = order
}

func (res *Transformer) TransformOrderCreate(item models.Order) {
	order := Order{}
	order.Id = item.Id
	order.TransactionNumber = item.TransactionNumber
	order.Media = item.Media
	order.IsMember = item.IsMember
	order.CustomerName = item.CustomerName
	order.CustomerEmail = item.CustomerEmail
	order.CustomerPhone = item.CustomerPhone
	order.CustomerAddress = item.CustomerAddress
	order.CreatedAt = item.CreatedAt
	order.UpdatedAt = item.UpdatedAt

	res.Data = order
}

func (res *Transformer) TransformOrderUpdate(item models.Order) {
	order := Order{}
	order.Id = item.Id
	order.TransactionNumber = item.TransactionNumber
	order.Media = item.Media
	order.IsMember = item.IsMember
	order.CustomerName = item.CustomerName
	order.CustomerEmail = item.CustomerEmail
	order.CustomerPhone = item.CustomerPhone
	order.CustomerAddress = item.CustomerAddress
	order.CreatedAt = item.CreatedAt
	order.UpdatedAt = item.UpdatedAt

	res.Data = order
}
