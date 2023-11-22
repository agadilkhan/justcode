package handler

import "lecture8/service"

type Handler struct {
	Service service.Service
}

func New(srvs service.Service) *Handler {
	return &Handler{srvs}
}
