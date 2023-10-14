package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	user := router.Group("/users")
	{
		user.POST("/register", h.Register)
		user.POST("/login", h.Login)
	}

	review := router.Group("/reviews")
	{
		review.Use(JWTVerify())
		review.GET("/", h.GetAllReviews)
		review.POST("/", h.CreateReview)
		review.GET("/:id", h.GetReview)
		review.DELETE("/:id", h.DeleteReview)
		review.PUT("/:id", h.UpdateReview)
	}

	return router
}
