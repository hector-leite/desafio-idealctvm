package routes

import (
	"github.com/hector-leite/desafio-idealctvm/application"
	"github.com/hector-leite/desafio-idealctvm/server/handler"
	"github.com/labstack/echo/v4"
)

func Register(app application.App, server *echo.Echo) {
	root := server.Group("/api")

	root.POST("/signup", handler.HandleSignUp(app.Services().User()))

	user := registerUser(app, root)
	registerUserFavoriteAsset(app, user)
}
