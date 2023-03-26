package handler

import (
	"net/http"

	"github.com/hector-leite/desafio-idealctvm/domain/contract"
	"github.com/hector-leite/desafio-idealctvm/server/viewmodel"
	"github.com/labstack/echo/v4"
)

func HandleSignUp(userService contract.UserService) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {

		var vm viewmodel.SignUpPlayerRequest
		err := ctx.Bind(&vm)
		if err != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, err)
		}

		err = userService.SignUp(vm.Parse())
		if err != nil {
			return respondGenericError(ctx, err)
		}

		return respondJSON(ctx, http.StatusNoContent, nil)
	}
}
