package handler

import (
	"github.com/gin-gonic/gin"
	"lecture10/models"
	"net/http"
)

func (h *Handler) Create(ctx *gin.Context) {
	var b models.Book

	err := ctx.ShouldBindJSON(&b)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, struct {
			Message string
		}{
			err.Error(),
		})

		return
	}

	id, err := h.Repo.Create(ctx, &b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, struct {
			Message string
		}{
			err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, struct {
		Data interface{}
	}{
		id,
	})
}
