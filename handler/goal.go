package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ehardi19/bitment-backend/domain"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllGoalByUser(ctx echo.Context) error {
	var response Response
	userID, err := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)
	}

	Goals, err := h.Services.GetAllGoalByUser(userID)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}
	response.Data = Goals
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) CreateGoal(ctx echo.Context) error {
	var response Response
	body := ctx.Request().Body
	var request domain.Goal
	err := json.NewDecoder(body).Decode(&request)
	if err != nil {
		response.Data = err.Error()
		response.Status = 401
		return ctx.JSON(401, response)
	}

	Goal, code, err := h.Services.CreateGoal(request)
	if err != nil {
		response.Data = err.Error()
		response.Status = code
		return ctx.JSON(code, response)
	}

	response.Data = Goal
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetGoalByID(ctx echo.Context) error {
	var response Response
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		return ctx.JSON(http.StatusBadRequest, response)
	}

	Goals, err := h.Services.GetGoalByID(id)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		return ctx.JSON(http.StatusInternalServerError, response)
	}
	response.Data = Goals
	response.Status = http.StatusOK

	return ctx.JSON(http.StatusOK, response)
}
