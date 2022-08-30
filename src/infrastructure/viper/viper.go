package viper

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func SetValuesOnEnvironment() {
	viper.AddConfigPath("./")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	viper.WatchConfig()
}

func PrintEnvs() {
	for k, m := range viper.GetViper().AllSettings() {
		log.Printf("KEY--> %s : VALUE -->%s", k, m)
	}
}
