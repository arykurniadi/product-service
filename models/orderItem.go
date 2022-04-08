package models

import "time"

type (
	OrderItem struct {
		Id        int       `json:"id" gorm:"primary_key,column:id"`
		OrderId   int       `json:"order_id" gorm:"column:order_id"`
		Name      string    `json:"name" gorm:"column:name"`
		Price     float64   `json:"price" gorm:"column:price"`
		CreatedAt time.Time `gorm:"column:created_at" json:"created_at" sql:"DEFAULT:current_timestamp"`
		UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" sql:"DEFAULT:current_timestamp"`

		Order Order `gorm:"association_foreignkey:OrderId; foreignkey:OrderId" json:"order"`
	}
)

func (OrderItem) TableName() string {
	return "order_items"
}
