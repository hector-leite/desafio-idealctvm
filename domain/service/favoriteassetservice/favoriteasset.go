package favoriteuserassetservice

import (
	"github.com/hector-leite/desafio-idealctvm/domain/contract"
	"github.com/hector-leite/desafio-idealctvm/domain/entity"
)

type FavoriteAssetService struct {
	favoriteAssetOrderHandler contract.FavoriteAssetOrderHandler
	data                      contract.DataManager
	yahooIntegration          contract.YahooFinance
}

func NewFavoriteUserAssetService(
	favoriteAssetOrderHandler contract.FavoriteAssetOrderHandler,
	data contract.DataManager,
	yahooIntegration contract.YahooFinance,
) contract.FavoriteAssetService {
	return FavoriteAssetService{
		favoriteAssetOrderHandler: favoriteAssetOrderHandler,
		data:                      data,
		yahooIntegration:          yahooIntegration,
	}
}

func (s FavoriteAssetService) UpdateFavoriteAssets(favoriteAssets []entity.FavoriteUserAsset) error {
	for _, favoriteAsset := range favoriteAssets {
		err := s.data.UserFavoriteAsset().Upsert(favoriteAsset)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s FavoriteAssetService) GetUserFavoriteAssets(userID uint, orderType string) ([]entity.Quote, error) {
	favoriteAssets, err := s.favoriteAssetOrderHandler.GetUserFavoriteAssets(userID, orderType)
	if err != nil {
		return []entity.Quote{}, err
	}

	return favoriteAssets, nil
}
