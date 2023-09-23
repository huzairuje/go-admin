package controllers

import (
	"github.com/labstack/echo/v4"
	"go-admin/infrastructure/session"
)

type GameController struct {
	BaseBackendController
}

func NewGameController() GameController {
	return GameController{
		BaseBackendController: BaseBackendController{
			Menu:        "Game",
			BreadCrumbs: []map[string]interface{}{},
		},
	}
}

func (c *GameController) Index(ctx echo.Context) error {
	breadCrumbs := map[string]interface{}{
		"menu": "Game",
		"link": "/check/game/home",
	}
	userInfo, err := session.Manager.Get(ctx, session.SessionId)
	if err != nil || userInfo == nil {
		session.SetFlashMessage(ctx, "session is expired! you must login first", "error", nil)
		return ctx.Redirect(302, "/check/auth/login")
	}
	return Render(ctx, "game", "index", c.Menu, session.GetFlashMessage(ctx),
		append(c.BreadCrumbs, breadCrumbs), userInfo)
}
