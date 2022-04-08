package customer

import "dbo.id/product-service/models"

type ICustomerRepository interface {
	GetListCustomer(int, int) ([]models.Customer, *models.Pagination, error)
	GetCustomerById(int) (models.Customer, error)
	GetCustomerByEmail(string) (models.Customer, error)
	Create(models.Customer) (models.Customer, error)
	Update(int, models.Customer) (models.Customer, error)
	Delete(int) error
}
