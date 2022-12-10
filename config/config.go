package config

import (
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App           AppAccount
	Routes        RoutesAccount
	Connection    ConnectionAccount
	Logger        LoggerAccount
	Authorization AuthorizationAccount
	CloudStorage  CloudStorageAccount
	Grafana       GrafanaAccount
}

type AppAccount struct {
	Name          string
	Endpoint      string
	Port          string
	Env           string
	SSL           string
	BodyLimit     int
	HexaSecretKey string
}

type RoutesAccount struct {
	Methods string
	Headers string
	Origins OriginsDetail
}

type OriginsDetail struct {
	IsDefault bool
	FeLocal   string
	FeDev     string
	FeProd    string
}

type ConnectionAccount struct {
	SimpleTransaction DatabaseAccount
	Redis             RedisAccount
}

type DatabaseAccount struct {
	DriverName      string
	DriverSource    string
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	MaxIdleConns    int
	ConnMaxIdleTime time.Duration
}

type RedisAccount struct {
	Host         string
	Password     string
	DB           int
	DefaultDB    int
	MinIdleConns int
	PoolSize     int
	PoolTimeout  time.Duration
}

type LoggerAccount struct {
	Logrus    LogrusAccount
	ZapLogger ZapLoggerAccount
}

type LogrusAccount struct {
	Level string
}

type ZapLoggerAccount struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

type AuthorizationAccount struct {
	JWT   JWTAccount
	Basic BasicAccount
}

type JWTAccount struct {
	AccessTokenSecretKey  string
	AccessTokenDuration   time.Duration
	RefreshTokenSecretKey string
	RefreshTokenDuration  time.Duration
}

type BasicAccount struct {
	ApiKey    string
	ApiSecret string
}

type CloudStorageAccount struct {
	GoogleStorage GoogleStorageAccount
}

type GoogleStorageAccount struct {
	ProjectID                string
	GoogleCredentialsFile    string
	GoogleCloudStorageBucket string
	GoogleCloudStorageURL    string
	AppName                  string
	DefaultMaxUploadSize     int
}

type GrafanaAccount struct {
	IsActive bool
	LokiURL  string
}

//=================================================================================================================

//* Init Config
func InitConfig(env string) *Config {
	configPath := GetConfigPath(env)

	confFile, err := LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	conf, err := ParseConfig(confFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}
	return conf
}

//* Get config path for local or docker
func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "./config/config-docker"
	}
	if configPath == "dev" {
		return "./config/config-dev"
	}
	if configPath == "staging" {
		return "./config/config-staging"
	}
	if configPath == "prod" {
		return "./config/config-prod"
	}
	return "./config/config-local"
}

//* Load config file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

//* Parse config ifle
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
