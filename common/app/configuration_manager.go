package app

import "product-app/common/postgre"

type ConfigurationManager struct {
	PostgreConfig postgre.Config
}

func NewConfigurationManager() *ConfigurationManager {
	postgreConfig := getPostgreConfig()
	return &ConfigurationManager{
		PostgreConfig: postgreConfig,
	}
}

func getPostgreConfig() postgre.Config {
	return postgre.Config{
		Host:                  "localhost",
		Port:                  "6432",
		DbName:                "productapp",
		UserName:              "postgres",
		Password:              "postgres",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	}
}
