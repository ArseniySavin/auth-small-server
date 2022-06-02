package config

import (
	"github.com/ArseniySavin/auth-small-server/internal/pkg/cache"
	"github.com/ArseniySavin/auth-small-server/internal/pkg/cache/etcd"
	catcher "github.com/ArseniySavin/catcher/pkg"
	"github.com/spf13/viper"
	"os"
)

var envFile = ".env"

// NewConfig -
func NewConfig() *Config {
	Conf := Config{
		Port:       viper.GetString(PORT),
		ETCD_ULRS:  viper.GetString(ETCD_ULRS),
		JWT_SECRET: viper.GetString(JWT_SECRET),
		DB: DB{
			Host:    viper.GetString(PG_HOST),
			Port:    viper.GetInt(PG_PORT),
			Name:    viper.GetString(PG_DB),
			Login:   viper.GetString(PG_LOGIN),
			Pass:    viper.GetString(PG_PASS),
			SslMode: viper.GetString(PG_SSL_MODE),
		},
	}

	return &Conf
}

// LoadDotEnv -
func LoadDotEnv() {
	if _, err := os.Stat(envFile); err != nil {
		if os.IsNotExist(err) {
			viper.AutomaticEnv()
			return
		}
		catcher.LogFatal(err)
	}

	viper.SetConfigFile(envFile)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		catcher.LogFatal(err)
	}
	return
}

func OpenCache(cfg *Config) {
	if err := cache.Open(etcd.Name, cfg.ETCD_ULRS); err != nil {
		catcher.LogFatal(err)
	}
}
