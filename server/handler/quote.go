package handler

import (
	"net/http"
	"strings"

	"github.com/hector-leite/desafio-idealctvm/domain/contract"
	"github.com/hector-leite/desafio-idealctvm/server/viewmodel"
	"github.com/labstack/echo/v4"
)

func HandleGetQuote(quoteService contract.QuoteService) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		assetValuesParam := ctx.QueryParam("asset")
		assetValues := strings.Split(assetValuesParam, ",")
		quotes, err := quoteService.GetQuotesByAssets(assetValues...)
		if err != nil {
			return respondGenericError(ctx, err)
		}

		return respondJSON(ctx, http.StatusOK, viewmodel.ParseGetQuotesByAssetsResponse(quotes))
	}
}
