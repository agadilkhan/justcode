package handler

import "github.com/gin-gonic/gin"

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	user := router.Group("/users")
	user.POST("/register", h.Register)
	user.POST("/login", h.Login)

	order := router.Group("/orders")
	order.Use(JWTVerify())
	order.GET("/", h.GetOrders)
	order.POST("/:id", h.CreateOrder)
	order.PUT("/:id", h.UpdateOrder)
	order.DELETE("/:id", h.DeleteOrder)
	return router
}
