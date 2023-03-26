package server

import (
	application "github.com/hector-leite/desafio-idealctvm/application"
	"github.com/hector-leite/desafio-idealctvm/server/routes"
	"github.com/hector-leite/desafio-idealctvm/server/serverconfig"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func Run(app application.App) {
	godotenv.Load()
	server := newServer(app)
	routes.Register(app, server)
	server.Logger.Fatal(server.Start(":" + app.Config.Server.Port))
}

func newServer(app application.App) *echo.Echo {
	e := echo.New()

	e.HideBanner = true

	e.Use(serverconfig.CorsMiddleware(app))
	if app.Config.App.Debug == "true" {
		e.Debug = true
	}
	return e
}
