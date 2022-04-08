package usecase

import (
	CustomerInterface "dbo.id/product-service/app/customer"
	"dbo.id/product-service/models"
	"dbo.id/product-service/requests"
	"github.com/gin-gonic/gin"
)

type CustomerUsecase struct {
	CustomerRepository CustomerInterface.ICustomerRepository
}

func NewCustomerUsecase(cs CustomerInterface.ICustomerRepository) CustomerInterface.ICustomerUsecase {
	return &CustomerUsecase{
		CustomerRepository: cs,
	}
}

func (cs *CustomerUsecase) GetListCustomer(c *gin.Context, page int, perPage int) (customers []models.Customer, pagination *models.Pagination, err error) {
	customers, pagination, err = cs.CustomerRepository.GetListCustomer(page, perPage)
	if err != nil {
		return nil, nil, err
	}

	return customers, pagination, nil
}

func (cs *CustomerUsecase) GetCustomerById(c *gin.Context, id int) (customer models.Customer, err error) {
	customer, err = cs.CustomerRepository.GetCustomerById(id)
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (cs *CustomerUsecase) Create(c *gin.Context, req requests.CustomerCreate) (customer models.Customer, err error) {
	customer = models.Customer{
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
	}

	customer, err = cs.CustomerRepository.Create(customer)
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (cs *CustomerUsecase) Update(c *gin.Context, id int, req requests.CustomerCreate) (customer models.Customer, err error) {
	customer = models.Customer{
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
	}

	customer, err = cs.CustomerRepository.Update(id, customer)
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (cs *CustomerUsecase) Delete(c *gin.Context, id int) (customer models.Customer, err error) {
	customer, err = cs.CustomerRepository.GetCustomerById(id)
	if err != nil {
		return customer, err
	}

	err = cs.CustomerRepository.Delete(id)
	if err != nil {
		return customer, err
	}

	return customer, nil
}
