package application

import (
	"github.com/hector-leite/desafio-idealctvm/domain/contract"
)

type AppService struct {
	userService               contract.UserService
	quoteService              contract.QuoteService
	favoriteAssetService      contract.FavoriteAssetService
	favoriteAssetOrderHandler contract.FavoriteAssetOrderHandler
}

func NewAppService(
	userService contract.UserService,
	quoteService contract.QuoteService,
	favoriteAssetService contract.FavoriteAssetService,
	favoriteAssetOrderHandler contract.FavoriteAssetOrderHandler,
) *AppService {
	return &AppService{
		userService:               userService,
		quoteService:              quoteService,
		favoriteAssetService:      favoriteAssetService,
		favoriteAssetOrderHandler: favoriteAssetOrderHandler,
	}
}

func (svc AppService) User() contract.UserService {
	return svc.userService
}

func (svc AppService) Quote() contract.QuoteService {
	return svc.quoteService
}

func (svc AppService) FavoriteAssetService() contract.FavoriteAssetService {
	return svc.favoriteAssetService
}

func (svc AppService) FavoriteAssetOrderHandler() contract.FavoriteAssetOrderHandler {
	return svc.favoriteAssetOrderHandler
}
