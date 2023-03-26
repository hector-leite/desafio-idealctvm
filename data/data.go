package data

import (
	"github.com/hector-leite/desafio-idealctvm/domain/contract"
	"github.com/hector-leite/desafio-idealctvm/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Conn struct {
	db *gorm.DB
}

func Connect(dsn string) (contract.DataManager, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	conn := new(Conn)
	conn.db = db

	conn.db.AutoMigrate(
		&entity.User{},
		&entity.FavoriteUserAsset{},
	)

	return conn, nil
}

func (c *Conn) User() contract.UserRepo {
	return newUserRepo(c.db)
}

func (c *Conn) UserFavoriteAsset() contract.UserFavoriteAssetRepo {
	return newUserFavoriteAsset(c.db)
}
