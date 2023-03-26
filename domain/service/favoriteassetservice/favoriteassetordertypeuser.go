package favoriteuserassetservice

import (
	"sort"

	"github.com/hector-leite/desafio-idealctvm/domain/contract"
	"github.com/hector-leite/desafio-idealctvm/domain/entity"
)

type FavoriteAssetOrderTypeUserService struct {
	data             contract.DataManager
	yahooIntegration contract.YahooFinance
}

func NewFavoriteAssetOrderTypeUserService(
	data contract.DataManager,
	yahooIntegration contract.YahooFinance,
) contract.FavoriteAssetOrderService {
	return FavoriteAssetOrderTypeUserService{
		data:             data,
		yahooIntegration: yahooIntegration,
	}
}

func (s FavoriteAssetOrderTypeUserService) GetUserFavoriteAssets(userID uint) ([]entity.Quote, error) {
	favoriteAssets, err := s.data.UserFavoriteAsset().GetUserFavoriteAssets(userID)
	if err != nil {
		return []entity.Quote{}, err
	}

	sort.Slice(favoriteAssets, func(i, j int) bool {
		return favoriteAssets[i].Position < favoriteAssets[j].Position
	})

	var favoriteAssetsString []string
	for _, favoriteAsset := range favoriteAssets {
		favoriteAssetsString = append(favoriteAssetsString, favoriteAsset.Symbol)
	}

	quotes, err := s.yahooIntegration.GetQuotesByAssets(favoriteAssetsString...)
	if err != nil {
		return []entity.Quote{}, err
	}

	orderedQuotes := make([]entity.Quote, len(quotes))
	for i, favoriteAsset := range favoriteAssets {
		for _, quote := range quotes {
			if quote.Symbol == favoriteAsset.Symbol {
				orderedQuotes[i] = quote
			}
		}
	}

	return quotes, nil
}
