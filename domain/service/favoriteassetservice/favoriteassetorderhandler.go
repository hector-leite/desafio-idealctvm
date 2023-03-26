package favoriteuserassetservice

import (
	"errors"

	"github.com/hector-leite/desafio-idealctvm/domain/contract"
	"github.com/hector-leite/desafio-idealctvm/domain/entity"
)

type favoriteAssetOrderHandler struct {
	favoriteAssetOrderService map[string]contract.FavoriteAssetOrderService
}

func NewFavoriteAssetHandler(
	favoriteAssetOrderService map[string]contract.FavoriteAssetOrderService,
) contract.FavoriteAssetOrderHandler {
	return favoriteAssetOrderHandler{
		favoriteAssetOrderService: favoriteAssetOrderService,
	}
}
func (s favoriteAssetOrderHandler) getOrderType(orderType string) (contract.FavoriteAssetOrderService, error) {
	if _, ok := s.favoriteAssetOrderService[orderType]; ok {
		return s.favoriteAssetOrderService[orderType], nil
	}

	return nil, errors.New("generic-error")
}

func (s favoriteAssetOrderHandler) GetUserFavoriteAssets(userID uint, orderType string) ([]entity.Quote, error) {
	orderFavoriteAssetService, err := s.getOrderType(orderType)
	if err != nil {
		return nil, err
	}

	favoriteAsset, err := orderFavoriteAssetService.GetUserFavoriteAssets(userID)
	if err != nil {
		return nil, err
	}
	return favoriteAsset, nil
}
