package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ehardi19/bitment-backend/domain"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetTotalInvestmentProjectionByUser(ctx echo.Context) error {
	var response Response
	userID, err := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)
	}

	Total, err := h.Services.GetTotalInvestmentProjectionByUser(userID)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}
	response.Data = Total
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetAllInvestmentByUser(ctx echo.Context) error {
	var response Response
	userID, err := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)
	}

	Investments, err := h.Services.GetAllInvestmentByUser(userID)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}
	response.Data = Investments
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) CreateInvestment(ctx echo.Context) error {
	var response Response
	body := ctx.Request().Body
	var request domain.Investment
	err := json.NewDecoder(body).Decode(&request)
	if err != nil {
		response.Data = err.Error()
		response.Status = 401
		return ctx.JSON(401, response)
	}

	Investment, code, err := h.Services.CreateInvestment(request)
	if err != nil {
		response.Data = err.Error()
		response.Status = code
		return ctx.JSON(code, response)
	}

	response.Data = Investment
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetInvestmentByID(ctx echo.Context) error {
	var response Response
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)
	}

	Investments, err := h.Services.GetInvestmentByID(id)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}
	response.Data = Investments
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}
