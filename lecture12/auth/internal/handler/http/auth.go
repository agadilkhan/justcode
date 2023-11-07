package http

import (
	"github.com/gin-gonic/gin"
	"lecture12/auth/internal/entity"
	"log"
	"net/http"
)

func (h *Handler) Register(ctx *gin.Context) {
	request := struct {
		Login    string
		Password string
	}{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("failed to Unmarshal err: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	id, err := h.service.Register(ctx, &entity.User{Login: request.Login, Password: request.Password})
	if err != nil {
		log.Printf("failed to Register err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, id)
}

func (h *Handler) Login(ctx *gin.Context) {
	request := struct {
		Login    string
		Password string
	}{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("failed to Unamrshal err: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	response, err := h.service.Login(ctx, request.Login, request.Password)
	if err != nil {
		log.Printf("failed to Login err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, response)
}
