package repository

import (
	"fmt"
	"strings"

	orderInterface "dbo.id/product-service/app/order"
	"dbo.id/product-service/models"
	"github.com/jinzhu/gorm"
)

type OrderRepository struct {
	ConnDB *gorm.DB
}

func NewOrderRepository(ConnDB *gorm.DB) orderInterface.IOrderRepository {
	return &OrderRepository{ConnDB}
}

func (m *OrderRepository) GetListOrder(page int, perPage int) (orders []models.Order, pagination *models.Pagination, err error) {
	tx := m.ConnDB

	if page < 1 {
		page = 1
	}

	if perPage < 1 {
		perPage = 25
	}

	offset := (page * perPage) - perPage

	_ = tx.Find(&orders).Error
	total := len(orders)

	if err := tx.Preload("Customer").Preload("OrderItems").Limit(perPage).Offset(offset).Find(&orders).Error; err != nil {
		return nil, nil, err
	}

	pagination = models.BuildPagination(total, page, perPage, len(orders))

	return orders, pagination, nil
}

func (m *OrderRepository) GetOrderById(id int) (order models.Order, err error) {
	tx := m.ConnDB.Begin()

	if err = tx.Preload("Customer").Preload("OrderItems").Where("id = ?", id).First(&order).Error; err != nil {
		return order, err
	}

	tx.Commit()
	return order, nil
}

func (m *OrderRepository) Create(value models.Order) (order models.Order, err error) {
	tx := m.ConnDB.Begin()

	if err := tx.Create(&value).Error; err != nil {
		tx.Rollback()

		return value, err
	}

	tx.Commit()
	return value, nil
}

func (m *OrderRepository) CreateOrderItem(orderItems []models.OrderItem) (err error) {
	tx := m.ConnDB.Begin()

	valueArgs := []interface{}{}
	valueStrings := []string{}

	for _, item := range orderItems {
		valueStrings = append(valueStrings, "(?, ?, ?, NOW(), NOW())")

		valueArgs = append(valueArgs, item.OrderId)
		valueArgs = append(valueArgs, item.Name)
		valueArgs = append(valueArgs, item.Price)
	}

	sql := "INSERT IGNORE INTO " + m.ConnDB.NewScope(models.OrderItem{}).TableName() + " (" +
		"order_id, " +
		"name, " +
		"price, " +
		"created_at, " +
		"updated_at" +
		") VALUES %s"

	sql = fmt.Sprintf(sql, strings.Join(valueStrings, ","))

	if err := tx.Exec(sql, valueArgs...).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
