package order

import "dbo.id/product-service/models"

type IOrderRepository interface {
	GetListOrder(int, int) ([]models.Order, *models.Pagination, error)
	GetOrderById(id int) (order models.Order, err error)
	Create(models.Order) (order models.Order, err error)
	Update(int, models.Order) (order models.Order, err error)
	CreateOrderItem([]models.OrderItem) (err error)
	DeleteOrderItem(int) (err error)
}
