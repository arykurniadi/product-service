package user

import (
	"dbo.id/product-service/models"
)

type IUserRepository interface {
	GetUserById(int) (models.User, error)
	GetListUser(int, int) ([]models.User, *models.Pagination, error)
	Create(user models.User) (models.User, error)
	Update(id int, value models.User) (models.User, error)
	Delete(id int) error
}
