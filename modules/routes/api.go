package routes

import (
	"github.com/labstack/echo/v4"
	"go-admin/boot"
)

func ApiRoutes(e *echo.Echo, setup boot.HandlerSetup) {
	apiRoutes := e.Group("/api")

	healthController := setup.HealthController
	healthRoutes := apiRoutes.Group("/health")
	healthRoutes.GET("/ping", healthController.Ping)
	healthRoutes.POST("/check", healthController.HealthCheckApi)

}
