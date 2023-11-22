package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) GetAllOrders(ctx *gin.Context) {
	orders, err := h.service.GetAllOrders(ctx)
	if err != nil {
		log.Printf("GetAllOrders err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func (h *Handler) GetOrderByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Printf("convert from string to int err: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}
	order, err := h.service.GetOrderByID(ctx, id)
	if err != nil {
		log.Printf("GetOrderByID err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, order)
}
