package config

import (
	"os"

	// "gitlab.diresoft.net/module/logger/logger"
	logger "github.com/sirupsen/logrus"
)

type Config struct {
	HttpListenAddress string
	ENV               string
	MQAddress         string
	Database
}

type Database struct {
	IP       string
	Password string
	User     string
}

var ENV_LOCAL = "local"
var ENV_DEV = "dev"
var ENV_STAGING = "staging"
var ENV_SIT = "sit"
var ENV_PROD = "prod"

var config *Config

func init() {
	setConfig()
}

func setConfig() {

	config = &Config{}

	// config.HttpListenAddress = os.Getenv("HttpListenAddress")
	// if config.HttpListenAddress == "" {
	// 	logger.Fatal("HttpListenAddress environment variable required but not set")
	// }

	// config.Database.IP = os.Getenv("DB_Address")
	// if config.Database.IP == "" {
	// 	logger.Fatal("DB_Address environment variable required but not set")
	// }

	// config.Database.Password = os.Getenv("DB_Password")

	// if config.Database.Password == "" {
	// 	logger.Fatal("DB_Password environment variable required but not set")
	// }

	// config.ENV = os.Getenv("ENV")
	// if config.ENV == "" {
	// 	logger.Fatal("ENV environment variable required but not set")
	// }

	// config.Database.User = "program"

	config.MQAddress = os.Getenv("MQ_Address")
	if config.MQAddress == "" {
		logger.Fatal("MQ_Address environment variable required but not set")
	}

}

func GetConfig() *Config {
	return config
}
