package config

import "os"

type Config struct {
	UserBalanceDB     string
	PasswordBalanceDB string
	NameDB            string
	ServerPort        string
}

func LoadConfig() *Config {
	config := Config{
		UserBalanceDB:     os.Getenv("avito_db_user"),
		PasswordBalanceDB: os.Getenv("avito_db_password"),
		NameDB:            os.Getenv("avito_db_name"),
		ServerPort:        "8080",
	}
	serverPort := os.Getenv("avito_server_port")
	if serverPort != "" {
		config.ServerPort = serverPort
	}
	return &config
}
