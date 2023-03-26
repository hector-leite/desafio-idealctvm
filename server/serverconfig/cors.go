package serverconfig

import (
	application "github.com/hector-leite/desafio-idealctvm/application"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CorsMiddleware(app application.App) echo.MiddlewareFunc {
	corsConfig := middleware.DefaultCORSConfig
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowOrigins = []string{app.Config.Server.WWWSources}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return middleware.CORSWithConfig(corsConfig)(next)(c)
		}
	}
}
