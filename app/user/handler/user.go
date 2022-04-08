package handler

import (
	"strconv"

	BaseHandler "dbo.id/product-service/app/api/handler"
	UserInterface "dbo.id/product-service/app/user"
	"dbo.id/product-service/requests"
	"dbo.id/product-service/transformers"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase UserInterface.IUserUsecase
}

func (a *UserHandler) GetListUser(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	perPage, _ := strconv.Atoi(c.Query("perPage"))

	users, pagination, err := a.UserUsecase.GetListUser(c, page, perPage)
	if err != nil {
		BaseHandler.RespondError(c, err.Error(), nil)
		return
	}

	res := new(transformers.CollectionPagingTransformer)
	res.TransformUserList(users, pagination)

	BaseHandler.RespondJSON(c, res, nil)
	return
}

func (a *UserHandler) GetUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := a.UserUsecase.GetUserById(c, id)
	if err != nil {
		BaseHandler.RespondError(c, err.Error(), nil)
		return
	}

	res := new(transformers.Transformer)
	res.TransformUserGetById(user)

	BaseHandler.RespondJSON(c, res, nil)
	return
}

func (a *UserHandler) Create(c *gin.Context) {
	req := requests.UserCreate{}
	err := c.ShouldBind(&req)
	if err != nil {
		BaseHandler.RespondError(c, err.Error(), nil)
		return
	}

	user, err := a.UserUsecase.Create(c, req)
	if err != nil {
		BaseHandler.RespondError(c, err.Error(), nil)
		return
	}

	res := new(transformers.Transformer)
	res.TransformUserCreate(user)

	BaseHandler.RespondJSON(c, res, nil)
	return
}

func (a *UserHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	req := requests.UserUpdate{}
	err := c.ShouldBind(&req)
	if err != nil {
		BaseHandler.RespondError(c, err.Error(), nil)
		return
	}

	user, err := a.UserUsecase.Update(c, id, req)
	if err != nil {
		BaseHandler.RespondError(c, err.Error(), nil)
		return
	}

	res := new(transformers.Transformer)
	res.TransformUserUpdate(user)

	BaseHandler.RespondJSON(c, res, nil)
	return
}

func (a *UserHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := a.UserUsecase.Delete(c, id)
	if err != nil {
		BaseHandler.RespondError(c, err.Error(), nil)
		return
	}

	res := new(transformers.Transformer)
	res.TransformUserDelete(user)

	BaseHandler.RespondJSON(c, res, nil)
	return
}
