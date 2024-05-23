package config

import "github.com/cristalhq/aconfig"

type Postgres struct {
	Host        string `env:"HOST"`
	Port        int    `env:"PORT"`
	User        string `env:"USER"`
	Password    string `env:"PASSWORD"`
	Database    string `env:"DATABASE"`
	SSLMode     string `env:"SSL_MODE"`
	SSLCertPath string `env:"SSL_CERT_PATH"`
}

type Redis struct {
	Host     string `env:"HOST"`
	Port     int    `env:"PORT"`
	Password string `env:"PASSWORD"`
	DB       int    `env:"DB"`
}

type Config struct {
	Debug   bool   `env:"DEBUG"`
	Address string `env:"ADDRESS"`

	Redis Redis    `env:"REDIS"`
	Pg    Postgres `env:"POSTGRES"`
}

func MustLoad() *Config {
	cfg := &Config{}

	err := aconfig.LoaderFor(cfg, aconfig.Config{
		EnvPrefix: "EDU",
	}).Load()

	if err != nil {
		panic(err)
	}

	return cfg
}
