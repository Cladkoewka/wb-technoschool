package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/spf13/viper"
)

type Config struct {
	Port uint16 `env:"PORT"`
	Host string `env:"HOST"`

	SQL DB `env-prefix:"POSTGRES_"`
	Kafka Kafka `env-prefix:"KAFKA_"`
}

type DB struct {
	Host     string `env:"HOST"`
	Port     uint16 `env:"PORT"`
	Username string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
	Name     string `env:"NAME"`
}

type Kafka struct {
	Broker  string `env:"BROKER"`
	Topic   string `env:"TOPIC"`
	GroupID string `env:"GROUP_ID"`
}


func LoadConfigFromFile(path string) (*Config, error) {
	config := new(Config)
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func MustLoadConfigFromFile(path string) *Config {
	config, err := LoadConfigFromFile(path)
	if err != nil {
		panic(err)
	}

	return config
}

func LoadConfigFromEnv() (*Config, error) {
	config := new(Config)
	err := cleanenv.ReadEnv(config)
	if err != nil {
		return nil, err
	}

	return config, err
}

func MustLoadConfigFromEnv() *Config {
	config, err := LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}

	return config
}