package favoriteuserassetservice

import (
	"sort"

	"github.com/hector-leite/desafio-idealctvm/domain/contract"
	"github.com/hector-leite/desafio-idealctvm/domain/entity"
)

type FavoriteAssetOrderTypeValueService struct {
	data             contract.DataManager
	yahooIntegration contract.YahooFinance
}

func NewFavoriteAssetOrderTypeValueService(
	data contract.DataManager,
	yahooIntegration contract.YahooFinance,
) contract.FavoriteAssetOrderService {
	return FavoriteAssetOrderTypeValueService{
		data:             data,
		yahooIntegration: yahooIntegration,
	}
}

func (s FavoriteAssetOrderTypeValueService) GetUserFavoriteAssets(userID uint) ([]entity.Quote, error) {
	favoriteAssets, err := s.data.UserFavoriteAsset().GetUserFavoriteAssets(userID)
	if err != nil {
		return []entity.Quote{}, err
	}

	var favoriteAssetsString []string
	for _, favoriteAsset := range favoriteAssets {
		favoriteAssetsString = append(favoriteAssetsString, favoriteAsset.Symbol)
	}

	quotes, err := s.yahooIntegration.GetQuotesByAssets(favoriteAssetsString...)
	if err != nil {
		return []entity.Quote{}, err
	}

	sort.Slice(quotes, func(i, j int) bool {
		return quotes[i].RegularMarketPrice > quotes[j].RegularMarketPrice
	})

	return quotes, nil
}
