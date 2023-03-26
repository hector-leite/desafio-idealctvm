package userservice

import (
	"github.com/hector-leite/desafio-idealctvm/domain/contract"
	"github.com/hector-leite/desafio-idealctvm/domain/entity"
)

type UserService struct {
	data contract.DataManager
}

func NewUserService(
	data contract.DataManager,
) contract.UserService {
	return UserService{
		data: data,
	}
}

func (s UserService) SignUp(user entity.User) error {
	_, err := s.data.User().SignUp(user)
	if err != nil {
		return err
	}

	return nil
}
