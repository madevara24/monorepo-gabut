package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

var (
	cfg  *Config
	once sync.Once
)

type Config struct {

	// METADATA
	AppName    string `mapstructure:"APP_NAME"`
	AppVersion string `mapstructure:"APP_VERSION"`

	// SERVER
	AllowedOrigins     string `mapstructure:"ALLOWED_ORIGINS"`
	ENV                string `mapstructure:"ENV"`
	Port               string `mapstructure:"PORT"`
	ServerReadTimeout  int    `mapstructure:"SERVER_READ_TIMEOUT"`
	ServerWriteTimeout int    `mapstructure:"SERVER_WRITE_TIMEOUT"`

	// DATABASE: POSTGRES
	DBHost        string `mapstructure:"DB_HOST"`
	DBMaxConn     int    `mapstructure:"DB_MAX_CONN"`
	DBMaxIdleConn int    `mapstructure:"DB_MAX_IDLE_CONN"`
	DBMaxTTLConn  int    `mapstructure:"DB_MAX_TTL_CONN"`
	DBName        string `mapstructure:"DB_NAME"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUsername    string `mapstructure:"DB_USERNAME"`
}

func Get() *Config {
	return cfg
}
func Load() {
	once.Do(func() {
		v := viper.New()
		v.AutomaticEnv()

		v.AddConfigPath(".")
		v.SetConfigType("env")
		v.SetConfigName(".env")
		v.AddConfigPath("/secrets")

		err := v.ReadInConfig()
		if err != nil {
			fmt.Println("config file not found: ", err)
		}

		config := new(Config)
		err = v.Unmarshal(config)
		if err != nil {
			panic(err)
		}

		cfg = config
	})
}
