package config

type Config struct {
	App          AppConfig
	Server       ServerConfig
	DB           DBConfig
	Cache        CacheConfig
	YahooFinance YahooFinanceConfig
}

type AppConfig struct {
	Debug string
}

type ServerConfig struct {
	Port       string
	WWWSources string
}

type DBConfig struct {
	Connection string
}

type CacheConfig struct {
	Host     string
	Port     int
	Prefix   string
	Password string
}

type YahooFinanceConfig struct {
	APIKey string
	APIURL string
}
