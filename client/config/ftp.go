package config

import "github.com/caarlos0/env/v6"

type Config struct {
	HOST string `env:"HOST" envDefault:"ftp.ncbi.nlm.nih.gov"`
	PORT int    `env:"PORT" envDefault:"21"`
	USER string `env:"USER" envDefault:"anonymous"`
	PASS string `env:"PASS" envDefault:""`
}

func FtpConfigs() (Config, error) {
	c := Config{}
	if err := env.Parse(&c); err != nil {
		return Config{}, err
	}
	return c, nil
}
