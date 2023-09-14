package middleware

import (
	"go-admin/config"
	"go-admin/infrastructure/session"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

const (
	oneWeekInSecond = 604800
)

func NewCookieStore() *sessions.CookieStore {
	authKeyConf := config.Conf.AuthKey
	encryptionKeyConf := config.Conf.EncryptionKey
	if authKeyConf == "" || encryptionKeyConf == "" {
		panic("auth key and or encryption key is either empty or not filled on config")
	}
	authKey := []byte(authKeyConf)
	encryptionKey := []byte(encryptionKeyConf)
	s := sessions.NewCookieStore(authKey, encryptionKey)
	maxAge := config.Conf.MaxAgeSession
	if maxAge == 0 {
		maxAge = oneWeekInSecond
	}
	s.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   maxAge,
		HttpOnly: true,
	}
	return s
}

func SessionMiddleware(s *session.ConfigSession) echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			result, err := s.Get(context, session.SessionId)
			if err != nil {
				return context.Redirect(302, "/check/auth/login")
			}
			if result == nil {
				return context.Redirect(302, "/check/auth/login")
			}
			return handlerFunc(context)

		}
	}
}
