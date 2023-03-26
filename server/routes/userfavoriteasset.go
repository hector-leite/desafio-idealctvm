package routes

import (
	"github.com/hector-leite/desafio-idealctvm/application"
	"github.com/hector-leite/desafio-idealctvm/server/handler"
	"github.com/labstack/echo/v4"
)

func registerUserFavoriteAsset(app application.App, user *echo.Group) {
	favorite := user.Group("/favorite")

	favorite.PUT("", handler.HandleUpdateFavoriteAsset(app.Services().FavoriteAssetService()))
	favorite.GET("", handler.HandleGetUserFavoriteAsset(app.Services().FavoriteAssetService()))
}
