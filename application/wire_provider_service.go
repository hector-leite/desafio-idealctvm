package application

import (
	"github.com/hector-leite/desafio-idealctvm/domain/contract"
	favoriteassetservice "github.com/hector-leite/desafio-idealctvm/domain/service/favoriteassetservice"
)

func ProvideFavoriteAssetHandler(
	data contract.DataManager,
	yahooFinance contract.YahooFinance,
) contract.FavoriteAssetOrderHandler {
	return favoriteassetservice.NewFavoriteAssetHandler(
		map[string]contract.FavoriteAssetOrderService{
			"alphabetical": favoriteassetservice.NewFavoriteAssetOrderTypeAlphabeticalService(
				data,
				yahooFinance,
			),
			"user": favoriteassetservice.NewFavoriteAssetOrderTypeUserService(
				data,
				yahooFinance,
			),
			"value": favoriteassetservice.NewFavoriteAssetOrderTypeValueService(
				data,
				yahooFinance,
			),
		},
	)
}
