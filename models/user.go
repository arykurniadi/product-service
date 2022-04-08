package models

type (
	User struct {
		Id       int    `json:"id" gorm:"primary_key,column:id"`
		Username string `json:"username" gorm:"column:username"`
		Password string `json:"password" gorm:"column:password"`
		Email    string `json:"email" gorm:"column:email"`
		Role     string `json:"role" gorm:"column:role"`
		Token    string `json:"token" gorm:"column:token"`
	}
)

func (User) TableName() string {
	return "users"
}
