package handler

import (
	"github.com/gin-gonic/gin"
	"lecture9/internal/entity"
	"net/http"
)

func (h *Handler) Register(ctx *gin.Context) {
	var u entity.User

	err := ctx.ShouldBindJSON(&u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Error{
			Message: err.Error(),
		})
		return
	}

	uId, err := h.Service.Register(ctx, &u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Error{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, OK{
		Message: "success",
		Data:    uId,
	})
}

func (h *Handler) Login(ctx *gin.Context) {
	loginRequest := struct {
		Username string
		Password string
	}{}

	err := ctx.ShouldBindJSON(loginRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Error{
			Message: err.Error(),
		})
		return
	}

	token, err := h.Service.Login(ctx, loginRequest.Username, loginRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Error{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, OK{
		Message: "success",
		Data:    token,
	})
}
