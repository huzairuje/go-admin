package routes

import (
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"go-admin/boot"
	"net/http"
)

func FrontendRoute(e *echo.Echo, setup boot.HandlerSetup) {
	e.Renderer = echoview.New(goview.Config{
		Root:         "views",
		Extension:    ".html",
		DisableCache: false,
	})

	e.GET("/", func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusTemporaryRedirect, "/check/auth/login")
	})

	frontendGroup := e.Group("/check")
	authController := setup.AuthController
	authGroup := frontendGroup.Group("/auth")
	authGroup.GET("/login", authController.Index)
	authGroup.POST("/login", authController.Login)
}
