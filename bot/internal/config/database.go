package config

import "fmt"

type PostgresConfig struct {
	Name     string `mapstructure:"name"`
	Type     string `mapstructure:"type"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

func (p *PostgresConfig) GetDSN() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s", p.Type, p.User, p.Password , p.Host, p.Port, p.Name)
}
