package routes

import (
	"github.com/hector-leite/desafio-idealctvm/application"
	"github.com/hector-leite/desafio-idealctvm/server/handler"
	"github.com/labstack/echo/v4"
)

func registerUser(app application.App, root *echo.Group) *echo.Group {
	user := root.Group("/user")

	user.GET("/quotes", handler.HandleGetQuote(app.Services().Quote()))

	return user
}
