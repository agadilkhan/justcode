package handler

import (
	"lecture12/orders/internal/config"
	"lecture12/orders/internal/service"
)

type Handler struct {
	service *service.Service
	cfg     *config.Config
}

func New(service *service.Service, cfg *config.Config) *Handler {
	return &Handler{
		service,
		cfg,
	}
}
