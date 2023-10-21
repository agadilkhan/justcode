package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) Get(ctx *gin.Context) {
	title := ctx.Query("title")

	book, err := h.BookCache.Get(ctx, title)
	if err != nil {
		return
	}

	if book == nil {
		book, err = h.Repo.Get(ctx, title)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, struct {
				Message string
			}{
				err.Error(),
			})
			return
		}
		err = h.BookCache.Set(ctx, title, book)
		if err != nil {
			log.Printf("could not cache book with title %s: %v", title, book)
		}
	}
	ctx.JSON(http.StatusOK, book)
}
