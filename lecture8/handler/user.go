package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"lecture8/entity"
	"lecture8/handler/dto"
	"net/http"
)

func (h *Handler) Register(ctx *gin.Context) {
	req := &dto.RegisterRequest{}

	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Error{
			Message: err.Error(),
		})
		return
	}

	userID, err := h.Service.Register(ctx, &entity.User{
		FirstName: req.FirstName, LastName: req.LastName, Username: req.Username, Password: req.Password},
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Error{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, OK{
		Message: "user registered success",
		Data:    userID,
	})
}

func (h *Handler) Login(ctx *gin.Context) {
	req := &dto.LoginRequest{}

	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Error{
			Message: err.Error(),
		})
		return
	}

	u, err := h.Service.Login(ctx, req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Error{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, OK{
		Message: "user sign in success",
		Data:    u.Token,
	})
}

const authUserID = "userID"

func getAuthUserID(ctx *gin.Context) (int, error) {
	userID, ok := ctx.MustGet(authUserID).(int)
	if !ok {
		ctx.JSON(http.StatusBadRequest, Error{
			Message: "can't get user id from auth",
		})
		return 0, errors.New("can't get user id from auth")
	}

	return userID, nil
}
