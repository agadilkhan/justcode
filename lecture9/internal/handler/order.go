package handler

import (
	"github.com/gin-gonic/gin"
	"lecture9/internal/entity"
	"net/http"
	"strconv"
)

func (h *Handler) GetOrders(ctx *gin.Context) {
	orders, err := h.Service.GetOrders(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Error{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, OK{
		Message: "success",
		Data:    orders,
	})
}

func (h *Handler) GetOrderByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Error{
			Message: err.Error(),
		})
		return
	}

	o, err := h.Service.GetOrderByID(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Error{
			err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, OK{
		Message: "success",
		Data:    o,
	})
}

func (h *Handler) CreateOrder(ctx *gin.Context) {
	var o entity.Order

	err := ctx.ShouldBindJSON(&o)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Error{
			Message: err.Error(),
		})
		return
	}

	oId, err := h.Service.CreateOrder(ctx, &o)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Error{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, OK{
		Message: "success",
		Data:    oId,
	})
}

func (h *Handler) DeleteOrder(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Error{
			Message: err.Error(),
		})
		return
	}

	oId, err := h.Service.DeleteOrder(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Error{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, OK{
		Message: "success",
		Data:    oId,
	})
}

func (h *Handler) UpdateOrder(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Error{
			Message: err.Error(),
		})
		return
	}

	var o entity.Order

	err = ctx.ShouldBindJSON(&o)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Error{
			Message: err.Error(),
		})
		return
	}

	o.ID = uint(id)

	newOrder, err := h.Service.UpdateOrder(ctx, &o)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Error{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, OK{
		Message: "success",
		Data:    newOrder,
	})
}
