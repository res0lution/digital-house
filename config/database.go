package config

import "github.com/res0lution/digital-house/utils"

type DatabaseConfig struct {
	Host string
	Name string
	Password string
	User string
	Port string
}

func NewDatabase() *DatabaseConfig {
	return &DatabaseConfig{
		Host: utils.GetIni("database", "HOST", "localhost"),
		Name: utils.GetIni("database", "DATABASE_NAME", "digital_house_db"),
		User: utils.GetIni("database", "DATABASE_USER", "postgres"),
		Password: utils.GetIni("database", "DATABASE_PASSWORD", "postgres"),
		Port: utils.GetIni("database", "PORT", "5432"),
	}
}