package data

import (
	"github.com/hector-leite/desafio-idealctvm/domain/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userFavoriteAsset struct {
	db *gorm.DB
}

func newUserFavoriteAsset(db *gorm.DB) userFavoriteAsset {
	return userFavoriteAsset{
		db: db,
	}
}

func (r userFavoriteAsset) Upsert(favoriteUserAsset entity.FavoriteUserAsset) error {
	result := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "symbol"}, {Name: "region"}},
		DoUpdates: clause.AssignmentColumns([]string{"favorite", "position"}),
	}).Create(&entity.FavoriteUserAsset{
		UserID:   favoriteUserAsset.UserID,
		Symbol:   favoriteUserAsset.Symbol,
		Region:   favoriteUserAsset.Region,
		Favorite: favoriteUserAsset.Favorite,
		Position: favoriteUserAsset.Position,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r userFavoriteAsset) GetUserFavoriteAssets(userID uint) ([]entity.FavoriteUserAsset, error) {
	var favoriteUserAsset []entity.FavoriteUserAsset

	result := r.db.Where("user_id = ? and favorite = true", userID).Find(&favoriteUserAsset)
	if result.Error != nil {
		return []entity.FavoriteUserAsset{}, result.Error
	}

	return favoriteUserAsset, nil
}
