package contract

import (
	"github.com/hector-leite/desafio-idealctvm/domain/entity"
)

type UserService interface {
	SignUp(user entity.User) error
}

type QuoteService interface {
	GetQuotesByAssets(assets ...string) ([]entity.Quote, error)
}

type FavoriteAssetService interface {
	UpdateFavoriteAssets(favoriteAssets []entity.FavoriteUserAsset) error
	GetUserFavoriteAssets(userID uint, orderType string) ([]entity.Quote, error)
}

type FavoriteAssetOrderHandler interface {
	GetUserFavoriteAssets(userID uint, orderType string) ([]entity.Quote, error)
}

type FavoriteAssetOrderService interface {
	GetUserFavoriteAssets(userID uint) ([]entity.Quote, error)
}
