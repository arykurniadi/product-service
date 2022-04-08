package repository

import (
	customerInterface "dbo.id/product-service/app/customer"
	"dbo.id/product-service/models"
	"github.com/jinzhu/gorm"
)

type CustomerRepository struct {
	ConnDB *gorm.DB
}

func NewCustomerRepository(ConnDB *gorm.DB) customerInterface.ICustomerRepository {
	return &CustomerRepository{ConnDB}
}

func (m *CustomerRepository) GetListCustomer(page int, perPage int) (customers []models.Customer, pagination *models.Pagination, err error) {
	tx := m.ConnDB

	if page < 1 {
		page = 1
	}

	if perPage < 1 {
		perPage = 25
	}

	offset := (page * perPage) - perPage

	_ = tx.Find(&customers).Error
	total := len(customers)

	if err := tx.Limit(perPage).Offset(offset).Find(&customers).Error; err != nil {
		return customers, nil, err
	}

	pagination = models.BuildPagination(total, page, perPage, len(customers))

	return customers, pagination, nil
}

func (m *CustomerRepository) GetCustomerById(id int) (customer models.Customer, err error) {
	tx := m.ConnDB.Begin()

	if err = tx.Where("id = ?", id).First(&customer).Error; err != nil {
		return customer, err
	}

	tx.Commit()
	return customer, nil
}

func (m *CustomerRepository) GetCustomerByEmail(email string) (customer models.Customer, err error) {
	tx := m.ConnDB.Begin()

	if err = tx.Where("email = ?", email).First(&customer).Error; err != nil {
		return customer, err
	}

	tx.Commit()
	return customer, nil
}

func (m *CustomerRepository) Create(value models.Customer) (customer models.Customer, err error) {
	tx := m.ConnDB.Begin()

	if err := tx.Create(&value).Error; err != nil {
		tx.Rollback()

		return value, err
	}

	tx.Commit()
	return value, nil
}

func (m *CustomerRepository) Update(id int, value models.Customer) (customer models.Customer, err error) {
	tx := m.ConnDB.Begin()

	if err := tx.Model(&customer).Where("id = ?", id).Updates(value).Error; err != nil {
		tx.Rollback()
		return customer, err
	}

	tx.Commit()

	return customer, nil
}

func (m *CustomerRepository) Delete(id int) (err error) {
	var customer models.Customer

	tx := m.ConnDB.Begin()

	if err := tx.Where("id = ?", id).Delete(&customer).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
