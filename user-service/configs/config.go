package configs

import (
	"bytes"
	_ "embed"
	"github.com/spf13/viper"
)

//go:embed config.yaml
var content []byte

type Mongo struct {
	MongoUser string
	MongoPass string
	MongoHost string
	MongoPort int
	MongoDB   string
	UserColl  string
}

type App struct {
	Debug           bool
	Host            string
	Port            int
	SecretKey       string
	TokenExpireHour int
}

type Config struct {
	Mongo *Mongo
	App   *App
}

func NewConfig() (*Config, error) {
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(bytes.NewBuffer(content)); err != nil {
		return nil, err
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
