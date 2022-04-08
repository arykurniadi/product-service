package handler

import (
	"strconv"

	BaseHandler "dbo.id/product-service/app/api/handler"
	CustomerInterface "dbo.id/product-service/app/customer"
	"dbo.id/product-service/requests"
	"dbo.id/product-service/transformers"
	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	CustomerUsercase CustomerInterface.ICustomerUsecase
}

func (a *CustomerHandler) GetListCustomer(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	perPage, _ := strconv.Atoi(c.Query("perPage"))

	customers, pagination, err := a.CustomerUsercase.GetListCustomer(c, page, perPage)
	if err != nil {
		BaseHandler.RespondError(c, err.Error(), nil)
		return
	}

	res := new(transformers.CollectionPagingTransformer)
	res.TransformCustomerList(customers, pagination)

	BaseHandler.RespondJSON(c, res, nil)
	return
}

func (a *CustomerHandler) GetCustomerById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	customer, err := a.CustomerUsercase.GetCustomerById(c, id)
	if err != nil {
		BaseHandler.RespondError(c, err.Error(), nil)
		return
	}

	res := new(transformers.Transformer)
	res.TransformCustomerGetById(customer)

	BaseHandler.RespondJSON(c, res, nil)
	return
}

func (a *CustomerHandler) Create(c *gin.Context) {
	req := requests.CustomerCreate{}
	err := c.ShouldBind(&req)
	if err != nil {
		BaseHandler.RespondError(c, err.Error(), nil)
		return
	}

	customer, err := a.CustomerUsercase.Create(c, req)
	if err != nil {
		BaseHandler.RespondError(c, err.Error(), nil)
		return
	}

	res := new(transformers.Transformer)
	res.TransformCustomerCreate(customer)

	BaseHandler.RespondJSON(c, res, nil)
	return
}

func (a *CustomerHandler) Update(c *gin.Context) {
	req := requests.CustomerCreate{}
	err := c.ShouldBind(&req)
	if err != nil {
		BaseHandler.RespondError(c, err.Error(), nil)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	customer, err := a.CustomerUsercase.Update(c, id, req)
	if err != nil {
		BaseHandler.RespondError(c, err.Error(), nil)
		return
	}

	res := new(transformers.Transformer)
	res.TransformCustomerUpdate(customer)

	BaseHandler.RespondJSON(c, res, nil)
	return
}

func (a *CustomerHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	customer, err := a.CustomerUsercase.Delete(c, id)
	if err != nil {
		BaseHandler.RespondError(c, err.Error(), nil)
		return
	}

	res := new(transformers.Transformer)
	res.TransformCustomerDelete(customer)

	BaseHandler.RespondJSON(c, res, nil)
	return
}
