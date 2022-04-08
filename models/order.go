package models

import "time"

type (
	Order struct {
		Id                int       `json:"id" gorm:"primary_key,column:id"`
		TransactionNumber string    `json:"transaction_number" gorm:"column:transaction_number"`
		Media             string    `json:"media" gorm:"column:media"`
		CustomerId        int       `json:"customer_id" gorm:"column:customer_id"`
		IsMember          int       `json:"is_member" gorm:"column:is_member"`
		CustomerName      string    `json:"customer_name" gorm:"column:customer_name"`
		CustomerEmail     string    `json:"customer_email" gorm:"column:customer_email"`
		CustomerPhone     string    `json:"customer_phone" gorm:"column:customer_phone"`
		CustomerAddress   string    `json:"customer_address" gorm:"column:customer_address"`
		CreatedAt         time.Time `gorm:"column:created_at" json:"created_at" sql:"DEFAULT:current_timestamp"`
		UpdatedAt         time.Time `gorm:"column:updated_at" json:"updated_at" sql:"DEFAULT:current_timestamp"`
		DeletedAt         time.Time `gorm:"column:deleted_at" json:"deleted_at" sql:"DEFAULT:current_timestamp"`

		Customer   Customer    `gorm:"association_foreignkey:Id; foreignkey:CustomerId" json:"customer"`
		OrderItems []OrderItem `gorm:"association_foreignkey:Id; foreignkey:OrderId" json:"order_items"`
	}
)

func (Order) TableName() string {
	return "orders"
}
