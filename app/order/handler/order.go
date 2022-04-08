package handler

import (
	"strconv"

	BaseHandler "dbo.id/product-service/app/api/handler"
	OrderInterface "dbo.id/product-service/app/order"
	"dbo.id/product-service/requests"
	"dbo.id/product-service/transformers"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	OrderUsecase OrderInterface.IOrderUsecase
}

func (od *OrderHandler) GetListOrder(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	perPage, _ := strconv.Atoi(c.Query("perPage"))

	orders, pagination, err := od.OrderUsecase.GetListOrder(c, page, perPage)
	if err != nil {
		BaseHandler.RespondError(c, err.Error(), nil)
		return
	}

	res := new(transformers.CollectionPagingTransformer)
	res.TransformOrderList(orders, pagination)

	BaseHandler.RespondJSON(c, res, nil)
	return
}

func (od *OrderHandler) GetOrderById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	order, err := od.OrderUsecase.GetOrderById(c, id)
	if err != nil {
		BaseHandler.RespondError(c, err.Error(), nil)
		return
	}

	res := new(transformers.Transformer)
	res.TransformOrderGetById(order)

	BaseHandler.RespondJSON(c, res, nil)
	return
}

func (od *OrderHandler) Create(c *gin.Context) {
	req := requests.OrderCreate{}
	err := c.ShouldBind(&req)
	if err != nil {
		BaseHandler.RespondError(c, err.Error(), nil)
		return
	}

	order, err := od.OrderUsecase.Create(c, req)
	if err != nil {
		BaseHandler.RespondError(c, err.Error(), nil)
		return
	}

	res := new(transformers.Transformer)
	res.TransformOrderCreate(order)

	BaseHandler.RespondJSON(c, res, nil)
	return
}

func (od *OrderHandler) Update(c *gin.Context) {
	req := requests.OrderCreate{}
	err := c.ShouldBind(&req)
	if err != nil {
		BaseHandler.RespondError(c, err.Error(), nil)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	order, err := od.OrderUsecase.Update(c, id, req)
	if err != nil {
		BaseHandler.RespondError(c, err.Error(), nil)
		return
	}

	res := new(transformers.Transformer)
	res.TransformOrderUpdate(order)

	BaseHandler.RespondJSON(c, res, nil)
	return
}
