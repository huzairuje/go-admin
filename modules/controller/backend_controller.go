package controllers

import (
	"net/http"

	"go-admin/infrastructure/session"

	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
)

type BaseBackendController struct {
	Title       string
	Menu        string
	BreadCrumbs []map[string]interface{}
}

func Render(ctx echo.Context, title string, view string, activeMenu string, flashMessage session.FlashMessage,
	breadcrumbs []map[string]interface{}, data interface{}) error {
	return echoview.Render(ctx, http.StatusOK, view, echo.Map{
		"title":        title,
		"activeMenu":   activeMenu,
		"breadCrumbs":  breadcrumbs,
		"ctx":          ctx,
		"flashMessage": flashMessage,
		"data":         data,
	})
}
