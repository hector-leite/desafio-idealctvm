package data

import (
	"github.com/hector-leite/desafio-idealctvm/domain/entity"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func newUserRepo(db *gorm.DB) userRepo {
	return userRepo{
		db: db,
	}
}

func (r userRepo) SignUp(user entity.User) (uint, error) {
	user.UUID = uuid.NewV4().String()
	result := r.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}
