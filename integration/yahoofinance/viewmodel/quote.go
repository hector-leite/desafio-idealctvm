package viewmodel

import "github.com/hector-leite/desafio-idealctvm/domain/entity"

type QuoteResponse struct {
	QuoteResponseResponse QuoteResponseResponse `json:"quoteResponse"`
}

type QuoteResponseResponse struct {
	QuoteResults []QuoteResult `json:"result"`
}

type QuoteResult struct {
	Symbol             string  `json:"symbol"`
	RegularMarketPrice float64 `json:"regularMarketPrice"`
	Currency           string  `json:"currency"`
	DisplayName        string  `json:"displayName"`
	Region             string  `json:"region"`
}

func (q QuoteResponse) Parse() []entity.Quote {
	quotes := make([]entity.Quote, len(q.QuoteResponseResponse.QuoteResults))
	for i, quote := range q.QuoteResponseResponse.QuoteResults {
		quotes[i].Symbol = quote.Symbol
		quotes[i].RegularMarketPrice = quote.RegularMarketPrice
		quotes[i].Currency = quote.Currency
		quotes[i].DisplayName = quote.DisplayName
		quotes[i].Region = quote.Region
	}

	return quotes
}
