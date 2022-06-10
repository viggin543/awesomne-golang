package models

import (
	"github.com/caarlos0/env/v6"
	"github.com/sirupsen/logrus"
)

type Config struct {
	MongoUri string `env:"MONGO_URI" envDefault:"mongodb://localhost:27017" json:"MONGO_URI"`
	USER     string `env:"MONGO_USER" envDefault:"root" json:"MONGO_USER"`
	PASS     string `env:"MONGO_PASS" envDefault:"root" json:"MONGO_PASS"`
	DBName   string `env:"MONGO_DB" envDefault:"test" json:"MONGO_DB"`
}

func ParseConfig() Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
	logrus.WithField("cfg", cfg).Info("parsed config")
	return cfg
}
