//go:build wireinject
// +build wireinject

package application

import (
	"github.com/google/wire"
	"github.com/hector-leite/desafio-idealctvm/domain/service/favoriteassetservice"
	"github.com/hector-leite/desafio-idealctvm/domain/service/quoteservice"
	"github.com/hector-leite/desafio-idealctvm/domain/service/userservice"
	"github.com/hector-leite/desafio-idealctvm/infra/config"
	"github.com/hector-leite/desafio-idealctvm/integration/yahoofinance"
)

func injectApp() App {
	panic(wire.Build(
		AppSet,

		wire.Struct(new(App),
			"Config",
			"DataManager",
			"CacheManager",
			"HTTPClient",
			"services",
			"integrations",
		),
	))
}

var AppSet = wire.NewSet(
	bindingSet,

	ProvideDB,
	ProvideCacheManager,
	ProvideConfig,
	ProvideHTTPClient,

	ProvideFavoriteAssetHandler,

	NewAppService,
	NewAppIntegration,
)

var bindingSet = wire.NewSet(
	configSet,
	serviceSet,
	integrationSet,
)

var configSet = wire.NewSet(
	wire.FieldsOf(new(*config.Config),
		"App",
		"Cache",
		"DB",
		"Server",
		"YahooFinance",
	),
)
var serviceSet = wire.NewSet(
	userservice.NewUserService,
	quoteservice.NewQuoteService,
	favoriteuserassetservice.NewFavoriteUserAssetService,
)

var integrationSet = wire.NewSet(
	yahoofinance.NewYahooFinanceIntegration,
)
