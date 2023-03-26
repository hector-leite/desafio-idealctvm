package entity

import "gorm.io/gorm"

type FavoriteUserAsset struct {
	gorm.Model
	Symbol   string `gorm:"not null; uniqueIndex:uidx_symbol_region_favorite_user:asc"`
	Region   string `gorm:"not null; uniqueIndex:uidx_symbol_region_favorite_user:asc"`
	Favorite bool   `gorm:"not null"`
	Position int    `gorm:"not null"`

	UserID uint `gorm:"not null; uniqueIndex:uidx_symbol_region_favorite_user:asc"`
}
