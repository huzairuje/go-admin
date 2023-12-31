package config

import (
	"time"
)

var (
	Conf Config
	Env  string

	EnvironmentDev  = "DEV"
	EnvironmentUAT  = "UAT"
	EnvironmentProd = "PROD"
	ListOfIsland    map[uint64]string

	searchPath = []string{
		"/etc/amartha/go-disb-platform",
		"$HOME/.go-disb-platform",
		".",
	}
	configDefaults = map[string]interface{}{
		"port":       1234,
		"logLevel":   "DEBUG",
		"logFormat":  "text",
		"signString": "supersecret",
	}
	configName = map[string]string{
		"local": "config.local",
		"dev":   "config.dev",
		"uat":   "config.uat",
		"prod":  "config.prod",
		"test":  "config.test",
	}
)

type Config struct {
	Env           string         `mapstructure:"env"`
	Port          int            `mapstructure:"port"`
	MaxAgeSession int            `mapstructure:"maxAgeSession"`
	AuthKey       string         `mapstructure:"authKey"`
	EncryptionKey string         `mapstructure:"encryptionKey"`
	LogLevel      string         `mapstructure:"logLevel"`
	LogMode       bool           `mapstructure:"logMode"`
	LogFormat     string         `mapstructure:"logFormat"`
	Postgres      PostgresConfig `mapstructure:"postgres"`
	Redis         RedisConfig    `mapstructure:"redis"`
}

type NewRelic struct {
	Name       string `mapstructure:"name"`
	LicenceKey string `mapstructure:"licenceKey"`
	Active     bool   `mapstructure:"active"`
}

type PSQL struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Schema   string `mapstructure:"schema"`
	DBName   string `mapstructure:"dbName"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

// PostgresConfig ...
type PostgresConfig struct {
	ConnMaxLifetime    int  `mapstructure:"connectTimeout"`
	MaxOpenConnections int  `mapstructure:"maxOpenConnections"`
	MaxIdleConnections int  `mapstructure:"maxIdleConnections"`
	Master             PSQL `mapstructure:"master"`
	Slave              PSQL `mapstructure:"slave"`
}

type RedisConfig struct {
	Host                      string        `mapstructure:"host"`
	Password                  string        `mapstructure:"password"`
	DB                        int           `mapstructure:"db"`
	Port                      int           `mapstructure:"port"`
	KeySubmitSosialisasiTtl   time.Duration `mapstructure:"keySubmitSosialisasiTtl"`
	ValueSubmitSosialisasiTtl time.Duration `mapstructure:"valueSubmitSosialisasiTtl"`
}
