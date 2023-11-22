package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	order := router.Group("/api/orders")
	{
		order.Use(JWTVerify(h.cfg))
		order.GET("/")
		order.GET("/:id")
	}

	return router
}
