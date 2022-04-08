package customer

import (
	"dbo.id/product-service/models"
	"dbo.id/product-service/requests"
	"github.com/gin-gonic/gin"
)

type ICustomerUsecase interface {
	GetListCustomer(c *gin.Context, page int, perPage int) ([]models.Customer, *models.Pagination, error)
	GetCustomerById(c *gin.Context, id int) (models.Customer, error)
	Create(c *gin.Context, req requests.CustomerCreate) (models.Customer, error)
	Update(c *gin.Context, id int, req requests.CustomerCreate) (models.Customer, error)
	Delete(c *gin.Context, id int) (models.Customer, error)
}
