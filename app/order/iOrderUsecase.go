package order

import (
	"dbo.id/product-service/models"
	"dbo.id/product-service/requests"
	"github.com/gin-gonic/gin"
)

type IOrderUsecase interface {
	GetListOrder(c *gin.Context, page int, perPage int) ([]models.Order, *models.Pagination, error)
	GetOrderById(c *gin.Context, id int) (models.Order, error)
	Create(c *gin.Context, req requests.OrderCreate) (models.Order, error)
	Update(c *gin.Context, id int, req requests.OrderCreate) (order models.Order, err error)
}
