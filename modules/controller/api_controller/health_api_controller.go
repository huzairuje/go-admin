package api_controller

import (
	"net/http"

	"go-admin/infrastructure/httplib"
	"go-admin/infrastructure/logger"
	"go-admin/modules/controller"
	"go-admin/modules/service"
	"go-admin/utils"

	"github.com/labstack/echo/v4"
)

type HealthController struct {
	controllers.Controller
	service service.HealthService
}

func NewHealthController(service service.HealthService) HealthController {
	return HealthController{
		service: service,
	}
}

func (h *HealthController) Ping(c echo.Context) error {
	return httplib.SetSuccessResponse(c, http.StatusOK, http.StatusText(http.StatusOK), "pong")
}

func (h *HealthController) HealthCheckApi(c echo.Context) error {
	logCtx := "handler.HealthCheckApi"
	ctx := c.Request().Context()
	err := h.service.CheckUpTime(ctx)
	if err != nil {
		logger.Error(ctx, utils.ErrorLogFormat, err.Error(), logCtx, "h.serviceHealth.CheckUpTime")
		return httplib.SetErrorResponse(c, http.StatusInternalServerError, "something went wrong")
	}
	return httplib.SetSuccessResponse(c, http.StatusOK, http.StatusText(http.StatusOK), "healthy")
}
