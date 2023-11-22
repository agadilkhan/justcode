package http

import "github.com/gin-gonic/gin"

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/api/auth/")
	{
		auth.POST("register", h.Register)
		auth.POST("login", h.Login)
	}

	return router
}
