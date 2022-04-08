package usecase

import (
	"fmt"
	"time"

	CustomerInterface "dbo.id/product-service/app/customer"
	OrderInterface "dbo.id/product-service/app/order"
	"dbo.id/product-service/models"
	"dbo.id/product-service/requests"
	"github.com/gin-gonic/gin"
)

type OrderUsecase struct {
	OrderRepository    OrderInterface.IOrderRepository
	CustomerRepository CustomerInterface.ICustomerRepository
}

func NewOrderUsecase(od OrderInterface.IOrderRepository, cr CustomerInterface.ICustomerRepository) OrderInterface.IOrderUsecase {
	return &OrderUsecase{
		OrderRepository:    od,
		CustomerRepository: cr,
	}
}

func (od *OrderUsecase) GetListOrder(c *gin.Context, page int, perPage int) (orders []models.Order, pagination *models.Pagination, err error) {
	orders, pagination, err = od.OrderRepository.GetListOrder(page, perPage)
	if err != nil {
		return nil, nil, err
	}

	return orders, pagination, err
}

func (od *OrderUsecase) GetOrderById(c *gin.Context, id int) (order models.Order, err error) {
	order, err = od.OrderRepository.GetOrderById(id)
	fmt.Println("GetOrderById", order)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (od *OrderUsecase) Create(c *gin.Context, req requests.OrderCreate) (order models.Order, err error) {
	var orderItems []models.OrderItem

	order.TransactionNumber = req.TransactionNumber
	order.Media = req.Media
	order.IsMember = 0

	if req.IsMember {
		customer, err := od.CustomerRepository.GetCustomerByEmail(req.Customer.Email)
		if err != nil {
			return order, err
		}

		order.CustomerId = customer.Id
		order.IsMember = 1
	} else {
		order.CustomerName = req.Customer.Name
		order.CustomerEmail = req.Customer.Email
		order.CustomerPhone = req.Customer.Phone
		order.CustomerAddress = req.Customer.Address
	}

	order.CreatedAt = time.Now()

	res, err := od.OrderRepository.Create(order)
	if err != nil {
		return res, err
	}

	for _, item := range req.OrderItems {
		orderItem := models.OrderItem{}
		orderItem.OrderId = res.Id
		orderItem.Name = item.Name
		orderItem.Price = float64(item.Price)

		orderItems = append(orderItems, orderItem)
	}

	err = od.OrderRepository.CreateOrderItem(orderItems)
	if err != nil {
		return order, err
	}

	return order, nil
}
