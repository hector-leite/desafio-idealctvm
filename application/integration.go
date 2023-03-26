package application

import "github.com/hector-leite/desafio-idealctvm/domain/contract"

type AppIntegration struct {
	yahooFinanceIntegration contract.YahooFinance
}

func NewAppIntegration(
	yahooFinanceIntegration contract.YahooFinance,
) *AppIntegration {
	return &AppIntegration{
		yahooFinanceIntegration: yahooFinanceIntegration,
	}
}

func (svc AppIntegration) YahooFinance() contract.YahooFinance {
	return svc.yahooFinanceIntegration
}
