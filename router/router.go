package router

import (
	"io"
	"net/http"
	"os"

	"go-admin/boot"
	"go-admin/infrastructure/httplib"
	"go-admin/infrastructure/session"
	"go-admin/infrastructure/validator"
	middlewareFunc "go-admin/modules/middleware"
	"go-admin/modules/routes"
	"go-admin/utils"

	gorillaCtx "github.com/gorilla/context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerRouter struct {
	Setup boot.HandlerSetup
}

func NewHandlerRouter(setup boot.HandlerSetup) HandlerRouter {
	return HandlerRouter{
		Setup: setup,
	}
}

func NotFoundHandler(c echo.Context) error {
	// Check if the request accepts JSON
	headerAccept := c.Request().Header.Get("Accept")
	headerContentType := c.Request().Header.Get("Content-Type")
	headersRule := []string{headerAccept, headerContentType}
	jsonTypeHeader := "application/json"
	if utils.Contains(headersRule, jsonTypeHeader) {
		// Response with JSON body
		return httplib.SetErrorResponse(c, http.StatusNotFound, "Not Found Routes")
	}
	// Response with HTML
	return c.Render(http.StatusNotFound, "auth/error.html", nil)
}

func (h HandlerRouter) RouterWithMiddleware() *echo.Echo {
	c := echo.New()
	c.Use(middleware.Logger())

	echo.NotFoundHandler = func(c echo.Context) error {
		return NotFoundHandler(c)
	}

	c.Static("/check/assets", "assets")

	c.Pre(middleware.RemoveTrailingSlash())

	//Validation
	c.Validator = validator.NewValidator()

	//Set Middleware
	c.Use(middleware.Recover())
	c.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	c.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: io.MultiWriter(os.Stdout),
	}))

	c.Use(echo.WrapMiddleware(gorillaCtx.ClearHandler))

	session.Manager = session.NewSessionManager(middlewareFunc.NewCookieStore())

	routes.ApiRoutes(c, h.Setup)
	routes.BackendRoute(c, h.Setup)
	routes.FrontendRoute(c, h.Setup)

	return c
}
