package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lecture8/entity"
	"lecture8/handler/dto"
	"net/http"
	"strconv"
)

func (h *Handler) CreateReview(ctx *gin.Context) {
	var req dto.CreateReviewRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Error{
			Message: err.Error(),
		})
		return
	}

	userId, err := getAuthUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Error{
			Message: err.Error(),
		})
	}

	id, err := h.Service.CreateReview(ctx, &entity.Review{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userId,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Error{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, OK{
		Message: "success",
		Data:    id,
	})
}

func (h *Handler) GetReview(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Error{
			Message: err.Error(),
		})
		return
	}

	review, err := h.Service.GetReview(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Error{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, OK{
		Message: "success",
		Data:    review,
	})
}

func (h *Handler) GetAllReviews(ctx *gin.Context) {
	reviews, err := h.Service.GetAllReviews(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Error{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, OK{
		Message: "success",
		Data:    reviews,
	})
}

func (h *Handler) DeleteReview(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Error{
			Message: err.Error(),
		})
		return
	}

	id, err = h.Service.DeleteReview(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Error{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, OK{
		Message: "success",
		Data:    id,
	})
}

func (h *Handler) UpdateReview(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Error{
			Message: err.Error(),
		})
		return
	}

	var req dto.UpdateReviewRequest

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Error{
			Message: err.Error(),
		})
		return
	}

	fmt.Println(req)

	r, err := h.Service.UpdateReview(ctx, &entity.Review{ID: id, Title: req.Title, Content: req.Content})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Error{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, OK{
		Message: "success",
		Data:    r,
	})
}
