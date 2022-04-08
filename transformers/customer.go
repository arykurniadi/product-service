package transformers

import "dbo.id/product-service/models"

type (
	Customer struct {
		Id      int    `json:"id"`
		Name    string `json:"name"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
		Address string `json:"address"`
	}
)

func (res *CollectionPagingTransformer) TransformCustomerList(arr []models.Customer, pagination *models.Pagination) {
	for _, item := range arr {
		customer := Customer{}
		customer.Id = item.Id
		customer.Name = item.Name
		customer.Email = item.Email
		customer.Phone = item.Phone
		customer.Address = item.Address

		res.Data = append(res.Data, customer)
	}

	res.Meta = pagination
}

func (res *Transformer) TransformCustomerGetById(item models.Customer) {
	customer := Customer{}
	customer.Id = item.Id
	customer.Name = item.Name
	customer.Email = item.Email
	customer.Phone = item.Phone
	customer.Address = item.Address

	res.Data = customer
}

func (res *Transformer) TransformCustomerCreate(item models.Customer) {
	customer := Customer{}
	customer.Name = item.Name
	customer.Email = item.Email
	customer.Phone = item.Phone
	customer.Address = item.Address

	res.Data = customer
}

func (res *Transformer) TransformCustomerUpdate(item models.Customer) {
	customer := Customer{}
	customer.Id = item.Id
	customer.Name = item.Name
	customer.Email = item.Email
	customer.Phone = item.Phone
	customer.Address = item.Address

	res.Data = customer
}

func (res *Transformer) TransformCustomerDelete(item models.Customer) {
	customer := Customer{}
	customer.Id = item.Id
	customer.Name = item.Name
	customer.Email = item.Email
	customer.Phone = item.Phone
	customer.Address = item.Address

	res.Data = customer
}
