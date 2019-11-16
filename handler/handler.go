package handler

import (
	"net/http"

	"github.com/ehardi19/bitment-backend/services"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Services *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{
		Services: services,
	}
}

func (h *Handler) HelloWorld(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"hello": "world",
	})
}

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}
