package viewmodel

import "github.com/hector-leite/desafio-idealctvm/domain/entity"

type SignUpPlayerRequest struct {
	Name            string `json:"name"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (p SignUpPlayerRequest) Parse() entity.User {
	return entity.User{
		Name: p.Name,
	}
}
