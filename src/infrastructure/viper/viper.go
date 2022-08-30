package viper

import (
	"regexp"

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
		if viper.GetBool("ENABLE_SECRET_VALIDATION") {
			if ok := IsSecret(m.(string)); !ok {
				log.Printf("KEY--> %s : VALUE -->%s", k, m)
			}
		} else {
			log.Printf("KEY--> %s : VALUE -->%s", k, m)
		}
	}
}

func IsSecret(value string) bool {
	var re = regexp.MustCompile(`^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$`)
	return re.MatchString(value)
}
