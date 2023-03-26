package quoteservice

import (
	"fmt"
	"strings"

	"github.com/hector-leite/desafio-idealctvm/domain/contract"
	"github.com/hector-leite/desafio-idealctvm/domain/entity"
)

type QuoteService struct {
	data             contract.DataManager
	cache            contract.CacheManager
	yahooIntegration contract.YahooFinance
}

func NewQuoteService(
	data contract.DataManager,
	cache contract.CacheManager,
	yahooIntegration contract.YahooFinance,
) contract.QuoteService {
	return QuoteService{
		data:             data,
		cache:            cache,
		yahooIntegration: yahooIntegration,
	}
}

func (s QuoteService) GetQuotesByAssets(assets ...string) ([]entity.Quote, error) {
	cacheKey := fmt.Sprintf("get-quote-by-asset:%s", strings.Join(assets, "-"))
	var quotes []entity.Quote
	err := s.cache.GetStruct(cacheKey, &quotes)
	if err != nil {
		quotes, err := s.yahooIntegration.GetQuotesByAssets(assets...)
		if err != nil {
			return []entity.Quote{}, err
		}
		s.cache.SetStruct(cacheKey, quotes)
		return quotes, nil
	}

	return quotes, nil
}
