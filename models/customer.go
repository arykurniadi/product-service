package models

type (
	Customer struct {
		Id      int    `json:"id" gorm:"primary_key,column:id"`
		Name    string `json:"name" gorm:"column:name"`
		Email   string `json:"email" gorm:"column:email"`
		Phone   string `json:"phone" gorm:"column:phone"`
		Address string `json:"address" gorm:"column:address"`
	}
)

func (Customer) TableName() string {
	return "customers"
}
