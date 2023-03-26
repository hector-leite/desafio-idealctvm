package contract

import "github.com/hector-leite/desafio-idealctvm/domain/entity"

type YahooFinance interface {
	GetQuotesByAssets(assets ...string) ([]entity.Quote, error)
}
