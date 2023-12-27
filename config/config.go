package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Database DBConf `yaml:"database"`
}

type DBConf struct {
	Host string `yaml:"host" env-default:"0.0.0.0"`
	Port string `yaml:"port" env-default:"3306"`
	User string `yaml:"user" env-default:"test"`
	Name string `yaml:"name" env-default:"test"`
	Pass string `yaml:"pass" env-default:"password"`
}

func New() (*Config, error) {
	conf := &Config{}
	// err := cleanenv.ReadConfig("./config.yaml", conf)
	// if err != nil {
	// 	return nil, err
	// }

	err := cleanenv.ReadEnv(conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
