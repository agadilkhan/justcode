package handler

import "lecture9/internal/service"

type Handler struct {
	Service service.Service
}

func New(srvs service.Service) Handler {
	return Handler{Service: srvs}
}
