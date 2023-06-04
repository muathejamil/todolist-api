package config

import (
	"errors"
	"fmt"
	"github.com/tkanos/gonfig"
)

const (
	DefaultEnv string = "dev"
)

// Configuration struct defines the configuration
type Configuration struct {
	DataBase DataBase
	Redis    Redis
	Server   Server
}

// GetConfiguration gets configuration
func GetConfiguration(params ...string) (Configuration, error) {
	// Define configuration
	configuration := Configuration{}
	// Default env is dev
	env := DefaultEnv
	// Check if we have params to change the env
	if len(params) > 0 {
		env = params[0]
	}
	// Config file name
	fileName := fmt.Sprintf("./config/%s_config.json", env)
	// Get config
	err := gonfig.GetConf(fileName, &configuration)
	// Handle error
	if err != nil {
		return Configuration{}, errors.New("failed to load configurations")
	}
	// No errors all good return configuration
	return configuration, nil
}
