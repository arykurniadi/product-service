package user

import (
	"dbo.id/product-service/models"
	"dbo.id/product-service/requests"
	"github.com/gin-gonic/gin"
)

type IUserUsecase interface {
	GetListUser(c *gin.Context, page int, perPage int) ([]models.User, *models.Pagination, error)
	GetUserById(c *gin.Context, id int) (models.User, error)
	Create(c *gin.Context, req requests.UserCreate) (*models.User, error)
	Update(c *gin.Context, id int, req requests.UserUpdate) (*models.User, error)
	Delete(c *gin.Context, id int) (models.User, error)
}
