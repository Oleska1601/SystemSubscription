package config

import "github.com/ilyakaznacheev/cleanenv"

type (
	Config struct {
		App
		HTTP
		Logger
		Web
		Postgres
	}

	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}

	HTTP struct {
		Port string `yaml:"port"`
	}

	Logger struct {
		Level string `yaml:"level"`
	}

	Web struct {
		Path string `yaml:"path"`
	}

	Postgres struct {
		PoolMax int    `yaml:"pool_max"`
		PgUrl   string `yaml:"pg_url"`
	}
)

func New() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig("./config/config.yml", &cfg)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil

}
