package config

import (
	"fmt"
	"time"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DB_USERNAME           string
	DB_PASSWORD           string
	DB_PORT               string
	DB_HOST               string
	DB_NAME               string
	DB_CONN_MAX_IDLE_TIME time.Duration
	DB_MAX_IDLE_CONNS     int
	DB_MAX_OPEN_CONNS     int

	MID_SERVER_KEY string
	MID_CLIENT_KEY string

	SECRET_KEY string
}

func GetConfig(params ...string) Configuration {
	configuration := Configuration{}
	env := "dev"
	if len(params) > 0 {
		env = params[0]
	}
	fileName := fmt.Sprintf("./config/%s_config.json", env)
	gonfig.GetConf(fileName, &configuration)
	return configuration
}
