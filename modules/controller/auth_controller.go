package controllers

import (
	log "github.com/sirupsen/logrus"
	"go-admin/modules/helper"
	"net/http"

	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"go-admin/infrastructure/session"
	"go-admin/modules/primitive/dto"
	"go-admin/modules/service"
)

type AuthController struct {
	Controller
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) AuthController {
	return AuthController{
		authService: authService,
	}
}

func (c *AuthController) Index(ctx echo.Context) error {
	return echoview.Render(ctx, http.StatusOK, "auth/login", echo.Map{
		"title":        "Login Page",
		"flashMessage": session.GetFlashMessage(ctx),
	})
}

func (c *AuthController) Login(ctx echo.Context) error {
	var loginDto dto.LoginDto
	if err := ctx.Bind(&loginDto); err != nil {
		session.SetFlashMessage(ctx, "something went wrong", "error", nil)
		return ctx.Redirect(302, "/check/auth/login")
	}
	if err := ctx.Validate(&loginDto); err != nil {
		session.SetFlashMessage(ctx, "email and password must filled", "error", nil)
		return ctx.Redirect(302, "/check/auth/login")
	}

	//Search Email
	data, err := c.authService.FindUserByEmail(loginDto.Email)
	if err != nil {
		log.Error(err)
		session.SetFlashMessage(ctx, "something went wrong", "error", nil)
		return ctx.Redirect(302, "/check/auth/login")
	}
	if !helper.CheckPasswordHash(loginDto.Password, data.Password) {
		session.SetFlashMessage(ctx, "wrong email or password", "error", nil)
		return ctx.Redirect(302, "/check/auth/login")
	}
	//Search Branch
	userInfo := session.UserInfo{
		UserID:     data.UserID,
		Name:       data.Name,
		Email:      data.Email,
		TypeUser:   data.TypeUser,
		UserRoleID: data.UserRoleID,
	}
	if errSetSession := session.Manager.Set(ctx, session.SessionId, &userInfo); errSetSession != nil {
		log.Error(errSetSession.Error())
		session.SetFlashMessage(ctx, "something went wrong, try again", "error", nil)
		return ctx.Redirect(302, "/check/auth/login")
	}
	session.SetFlashMessage(ctx, "login success", "success", nil)

	return ctx.Redirect(302, "/check/admin/home")
}

func (c *AuthController) Logout(ctx echo.Context) error {
	err := session.Manager.Delete(ctx, session.SessionId)
	if err != nil {
		log.Error(err.Error())
		session.SetFlashMessage(ctx, "something went wrong, try again", "error", nil)
		return ctx.Redirect(302, "/check/admin/home")
	}
	session.SetFlashMessage(ctx, "logout success", "success", nil)
	return ctx.Redirect(http.StatusFound, "/check/auth/login")
}
