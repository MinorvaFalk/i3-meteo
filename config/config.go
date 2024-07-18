package config

import (
	"log"
	"os"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type config struct {
	Env      string `mapstructure:"ENV"`
	Port     string `mapstructure:"PORT"`
	Dsn      string `mapstructure:"DSN"`
	MeteoUrl string `mapstructure:"METEOSOURCE_URL"`
	MeteoKey string `mapstructure:"METEOSOURCE_API_KEY"`
}

var cfg *config

func InitConfig(filenames ...string) {
	vi := viper.NewWithOptions()

	if len(filenames) > 0 {
		if _, err := os.Stat(filenames[0]); err != nil {
			log.Fatal(err)
		}

		vi.SetConfigFile(filenames[0])
	} else {
		if _, err := os.Stat(".env"); err == nil {
			vi.SetConfigFile(".env")
		}
	}

	setDefaultValues(vi)
	vi.AutomaticEnv()

	if vi.ConfigFileUsed() != "" {
		if err := vi.ReadInConfig(); err != nil {
			log.Fatal(err)
		}
	} else {
		var out map[string]any
		if err := mapstructure.Decode(config{}, &out); err != nil {
			log.Fatal(err)
		}

		for key := range out {
			vi.MustBindEnv(key)
		}
	}

	if err := vi.Unmarshal(&cfg); err != nil {
		log.Fatal(err)
	}
}

func ReadConfig(filenames ...string) *config {
	if cfg == nil {
		InitConfig(filenames...)
	}

	return cfg
}

func setDefaultValues(vi *viper.Viper) {
	vi.SetDefault("ENV", "development")
	vi.SetDefault("PORT", "8080")
	vi.SetDefault("METEOSOURCE_URL", "https://www.meteosource.com/api/v1/free")
}
