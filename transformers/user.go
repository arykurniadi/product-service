package transformers

import (
	"dbo.id/product-service/models"
)

type (
	User struct {
		Id       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Role     string `json:"role"`
		Token    string `json:"token,omitempty"`
	}
)

func (res *CollectionPagingTransformer) TransformUserList(arrUser []models.User, pagination *models.Pagination) {
	for _, item := range arrUser {
		user := User{}
		user.Id = item.Id
		user.Username = item.Username
		user.Email = item.Email
		user.Role = item.Role

		res.Data = append(res.Data, user)
	}
	res.Meta = pagination
}

func (res *Transformer) TransformUserGetById(item models.User) {
	user := User{}
	user.Id = item.Id
	user.Username = item.Username
	user.Email = item.Email
	user.Role = item.Role
	user.Token = item.Token

	res.Data = user
}

func (res *Transformer) TransformUserCreate(item *models.User) {
	user := User{}
	user.Id = item.Id
	user.Username = item.Username
	user.Email = item.Email
	user.Role = item.Role
	user.Token = item.Token

	res.Data = user
}

func (res *Transformer) TransformUserUpdate(item *models.User) {
	user := User{}
	user.Id = item.Id
	user.Username = item.Username
	user.Email = item.Email
	user.Role = item.Role

	res.Data = user
}

func (res *Transformer) TransformUserDelete(item models.User) {
	user := User{}
	user.Id = item.Id
	user.Username = item.Username
	user.Email = item.Email
	user.Role = item.Role

	res.Data = user
}
