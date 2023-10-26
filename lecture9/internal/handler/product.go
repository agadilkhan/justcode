package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetProducts(ctx *gin.Context) {
	products, err := h.Service.GetProducts(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Error{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, OK{
		Message: "success",
		Data:    products,
	})
}
