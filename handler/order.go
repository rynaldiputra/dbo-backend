package handler

import (
	"dbo-be/helper"
	"dbo-be/order"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	order order.Service
}

func NewOrderHandler(order order.Service) *OrderHandler {
	return &OrderHandler{order}
}

func (o *OrderHandler) Search(c *gin.Context) {
	var input order.OrderSearchInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.ValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.JsonResponse("Gagal menyimpan data order", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	order, err := o.order.SearchOrders(input)

	if err != nil {
		response := helper.JsonResponse("Gagal menyimpan data order", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.JsonResponse("data order berhasil diambil", http.StatusOK, "success", order)

	c.JSON(http.StatusOK, response)
}

func (o *OrderHandler) Get(c *gin.Context) {
	order, err := o.order.GetOrders()

	if err != nil {
		response := helper.JsonResponse("Gagal mendapatkan data order", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.JsonResponse("data order berhasil diambil", http.StatusOK, "success", order)

	c.JSON(http.StatusOK, response)
}

func (o *OrderHandler) Find(c *gin.Context) {
	findOrder, err := o.order.GetOrderById(c.Param("id"))

	if err != nil {
		response := helper.JsonResponse("Gagal mendapatkan data order", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	msgFormatResp := order.FormatOrder(findOrder)
	response := helper.JsonResponse("data order berhasil diambil", http.StatusOK, "success", msgFormatResp)

	c.JSON(http.StatusOK, response)
}

func (o *OrderHandler) Create(c *gin.Context) {
	var input order.OrderInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.ValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.JsonResponse("Gagal menyimpan data order", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newOrder, err := o.order.CreateOrder(input)

	if err != nil {
		response := helper.JsonResponse("Gagal menyimpan data order", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	msgFormatResp := order.FormatOrder(newOrder)
	response := helper.JsonResponse("Penyimpanan data order berhasil", http.StatusOK, "success", msgFormatResp)

	c.JSON(http.StatusOK, response)
}

func (o *OrderHandler) Edit(c *gin.Context) {
	var input order.OrderInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.ValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.JsonResponse("Gagal menyimpan data order", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	editOrder, err := o.order.EditOrder(c.Param("id"), input)

	if err != nil {
		response := helper.JsonResponse("Gagal menyimpan data order", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	msgFormatResp := order.FormatOrder(editOrder)
	response := helper.JsonResponse("Penyimpanan data order berhasil", http.StatusOK, "success", msgFormatResp)

	c.JSON(http.StatusOK, response)
}

func (o *OrderHandler) Delete(c *gin.Context) {
	deleteOrder, err := o.order.DeleteOrder(c.Param("id"))

	if err != nil {
		response := helper.JsonResponse("Gagal menyimpan data order", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	msgFormatResp := order.FormatOrder(deleteOrder)
	response := helper.JsonResponse("Penghapusan data berhasil", http.StatusOK, "success", msgFormatResp)

	c.JSON(http.StatusOK, response)
}
