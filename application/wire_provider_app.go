package application

import (
	"net/http"
	"time"

	data "github.com/hector-leite/desafio-idealctvm/data"
	"github.com/hector-leite/desafio-idealctvm/domain/contract"
	"github.com/hector-leite/desafio-idealctvm/infra/cache"
	"github.com/hector-leite/desafio-idealctvm/infra/config"
	infrahttp "github.com/hector-leite/desafio-idealctvm/infra/http"
)

func ProvideDB(cfg config.DBConfig) contract.DataManager {
	db, err := data.Connect(cfg.Connection)
	if err != nil {
		panic(err)
	}
	return db
}

func ProvideCacheManager(cfg config.CacheConfig) contract.CacheManager {
	return cache.Connect(
		cfg.Host,
		cfg.Port,
		0,
		cfg.Password,
		cfg.Prefix,
		time.Duration(time.Duration(24)*time.Hour),
	)
}

func ProvideConfig() *config.Config {
	return config.Read()
}

func ProvideHTTPClient() contract.HTTPClient {
	return infrahttp.New(new(http.Client))
}
