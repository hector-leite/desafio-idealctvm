package viewmodel

import "github.com/hector-leite/desafio-idealctvm/domain/entity"

type UserFavoritesAssetsRequest []UserFavoriteAssetsRequest

type UserFavoriteAssetsRequest struct {
	Symbol   string `json:"symbol"`
	Region   string `json:"region"`
	Favorite bool   `json:"favorite"`
	Position int    `json:"position"`
}

func (vm UserFavoritesAssetsRequest) Parse() []entity.FavoriteUserAsset {
	userFavoritesAsset := make([]entity.FavoriteUserAsset, len(vm))
	for i, userFavoriteAsset := range vm {
		userFavoritesAsset[i].Symbol = userFavoriteAsset.Symbol
		userFavoritesAsset[i].Region = userFavoriteAsset.Region
		userFavoritesAsset[i].Favorite = userFavoriteAsset.Favorite
		userFavoritesAsset[i].Position = userFavoriteAsset.Position
		userFavoritesAsset[i].UserID = uint(1)

	}

	return userFavoritesAsset
}

type GetQuoteByAssetResponse struct {
	Symbol             string  `json:"symbol"`
	RegularMarketPrice float64 `json:"regularMarketPrice"`
	Currency           string  `json:"currency"`
	DisplayName        string  `json:"displayName"`
	Region             string  `json:"region"`
}

type GetQuotesByAssetsResponse struct {
	Result []GetQuoteByAssetResponse `json:"result"`
}

func ParseGetQuotesByAssetsResponse(quotes []entity.Quote) (response GetQuotesByAssetsResponse) {
	result := make([]GetQuoteByAssetResponse, len(quotes))
	for i, quote := range quotes {
		result[i].Symbol = quote.Symbol
		result[i].RegularMarketPrice = quote.RegularMarketPrice
		result[i].Currency = quote.Currency
		result[i].DisplayName = quote.DisplayName
		result[i].Region = quote.Region
	}
	response.Result = result
	return response
}
