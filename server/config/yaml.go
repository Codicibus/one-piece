package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	JWT struct {
		SecretKey      string `yaml:"secret_key"`
		ExpireTime int64    `yaml:"expire_time"`
		BufferTime int64 `yaml:"buffer_time"`
	} `yaml:"jwt"`
	DB struct {
		Type string `yaml:"type"`
	} `yaml:"db"`
	Postgres struct {
		DBPort     int    `yaml:"db_port"`
		DBHost     string `yaml:"db_host"`
		DBName     string `yaml:"db_name"`
		DBUsername string `yaml:"db_username"`
		DBPassword string `yaml:"db_password"`
		SSLMode    string   `yaml:"ssl_mode"`
		Timezone   string `yaml:"timezone"`
	} `yaml:"postgres"`
}

func ReadYAMLConfig() *Config {
	data, err := ioutil.ReadFile("cmd/opiece.yaml")
	if err != nil {
		panic(err)
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	return &config
}
