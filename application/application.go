package application

import (
	"github.com/hector-leite/desafio-idealctvm/domain/contract"
	"github.com/hector-leite/desafio-idealctvm/infra/config"
)

type App struct {
	Config       *config.Config
	DataManager  contract.DataManager
	CacheManager contract.CacheManager
	HTTPClient   contract.HTTPClient
	services     *AppService
	integrations *AppIntegration
}

func BuildApp() App {
	return injectApp()
}

func (app App) Services() *AppService {
	return app.services
}

func (app App) Integrations() *AppIntegration {
	return app.integrations
}
