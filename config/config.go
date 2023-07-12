package config

import (
	"github/yogabagas/print-in/pkg/config"
	"log"
	"os"
)

var GlobalCfg *Config

type (
	Config struct {
		App         App    `json:"app"`
		DB          DB     `json:"db"`
		PasswordAlg string `json:"password_alg"`
	}

	App struct {
		Name         string `json:"name"`
		Host         string `json:"host"`
		Port         string `json:"port"`
		ReadTimeout  int    `json:"read_timeout"`
		WriteTimeout int    `json:"write_timeout"`
	}

	DB struct {
		SQL struct {
			User     string `json:"user"`
			Password string `json:"password"`
			Host     string `json:"host"`
			Schema   string `json:"schema"`
		} `json:"sql"`
	}
)

func LoadConfig(path string) interface{} {

	if GlobalCfg == nil {
		err := config.ReadModuleConfig(
			&config.Cfg{
				Target: &GlobalCfg,
				Path:   path,
				Module: "config",
				Env:    os.Getenv("APP_ENV"),
			})
		if err != nil {
			log.Fatalln("can't load file config", err)
		}
	}
	return GlobalCfg
}
