package store

import "fmt"

type Config struct {
	Host     string `toml:"db_host"`
	Port     string `toml:"db_port"`
	Login    string `toml:"db_login"`
	Password string `toml:"db_password"`
	Database string `toml:"db_name"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) getDbUrl() string {
	return fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.Login, c.Password, c.Database)
}
