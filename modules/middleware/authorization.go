package middleware

import (
	"net/http"
	"strings"

	"go-admin/infrastructure/session"
	models "go-admin/modules/primitive/model"
	"go-admin/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type AuthorizationConfig struct {
	Skipper middleware.Skipper
}

func NewAuthorizationMiddleware() *AuthorizationConfig {
	return &AuthorizationConfig{
		Skipper: func(context echo.Context) bool {
			return false
		},
	}
}

func (m *AuthorizationConfig) AuthorizationMiddleware(menu []models.Menu, prefix string) echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			if m.Skipper(context) {
				return handlerFunc(context)
			}

			//check role
			result, err := session.Manager.Get(context, session.SessionId)
			if err != nil {
				return echo.NewHTTPError(404, "you must login before access this resource")
			}
			userInfo := result.(session.UserInfo)

			for _, v := range menu {
				if v.Prefix == prefix {
					if !utils.ItemExists(strings.Split(v.Role, ","), userInfo.UserRoleID) {
						return echo.NewHTTPError(http.StatusForbidden, "this user role don't have permission to access this resource ")
					} else {
						return handlerFunc(context)
					}
				}
			}
			return echo.NewHTTPError(500, "error get user from context ")
		}
	}
}
