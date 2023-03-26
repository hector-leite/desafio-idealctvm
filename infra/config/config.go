package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Read() *Config {
	godotenv.Load(".env")

	var config Config
	config.Cache.Port, _ = strconv.Atoi(os.Getenv("REDIS_PORT"))
	config.Cache.Host = os.Getenv("REDIS_HOSTNAME")
	config.Cache.Prefix = os.Getenv("REDIS_PREFIX")
	config.Cache.Password = os.Getenv("REDIS_PASSWORD")

	config.DB.Connection = os.Getenv("DATABASE_CONNECTION")

	config.App.Debug = os.Getenv("ENV_DEBUG")

	config.Server.WWWSources = os.Getenv("WWW_SOURCES")
	config.Server.Port = os.Getenv("PORT")

	config.YahooFinance.APIKey = os.Getenv("YAHOO_FINANCE_API_KEY")
	config.YahooFinance.APIURL = os.Getenv("YAHOO_FINANCE_API_URL")

	return &config
}
