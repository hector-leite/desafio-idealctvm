package handler

import (
	"net/http"

	"github.com/hector-leite/desafio-idealctvm/domain/contract"
	"github.com/hector-leite/desafio-idealctvm/server/viewmodel"
	"github.com/labstack/echo/v4"
)

func HandleGetUserFavoriteAsset(favoriteAssetService contract.FavoriteAssetService) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {

		orderParam := ctx.QueryParam("order")

		userID := uint(1)
		userFavoriteAssets, err := favoriteAssetService.GetUserFavoriteAssets(userID, orderParam)
		if err != nil {
			return respondGenericError(ctx, err)
		}

		return respondJSON(ctx, http.StatusOK, viewmodel.ParseGetQuotesByAssetsResponse(userFavoriteAssets))
	}
}

func HandleUpdateFavoriteAsset(favoriteAssetService contract.FavoriteAssetService) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {

		var vm viewmodel.UserFavoritesAssetsRequest
		err := ctx.Bind(&vm)
		if err != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, err)
		}

		err = favoriteAssetService.UpdateFavoriteAssets(vm.Parse())
		if err != nil {
			return respondGenericError(ctx, err)
		}

		return respondJSON(ctx, http.StatusNoContent, nil)
	}
}
