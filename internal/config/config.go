package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	PostgresURL string `mapstructure:"POSTGRES_URL"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() error {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("не удалось прочитать параметры: %w", err)
	}
	if err := viper.Unmarshal(c); err != nil {
		return fmt.Errorf("не удалось распарсить параметры: %w", err)
	}
	return nil
}

