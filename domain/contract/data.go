package contract

import (
	"github.com/hector-leite/desafio-idealctvm/domain/entity"
)

type DataManager interface {
	RepoManager
}

type RepoManager interface {
	User() UserRepo
	UserFavoriteAsset() UserFavoriteAssetRepo
}

type UserRepo interface {
	SignUp(user entity.User) (uint, error)
}

type UserFavoriteAssetRepo interface {
	Upsert(favoriteUserAsset entity.FavoriteUserAsset) error
	GetUserFavoriteAssets(userID uint) ([]entity.FavoriteUserAsset, error)
}
