package config

import (
	"fmt"
	"gotify/util"
	"os"
	"sync"
)

type ConfigManagerType string
const (
	CONFIG_NAIVE   ConfigManagerType = "naive"
	CONFIG_UNKNOWN ConfigManagerType = "unknown"
)

var (
	configManager ConfigManager
	configManagerOnce sync.Once
)

func Config() ConfigManager {
	configManagerOnce.Do(initialise)
	if configManager == nil {
		initialise()
	}
	return configManager
}

func initialise() {
	cm, err := ConfigFactory().Build()
	if err != nil {
		util.DoError(fmt.Errorf("ERROR|config/config.initialise()|failed to initialise config factory|%s", err.Error()))
		cm = &unknownConfigManager{}
	}
	configManager = cm
}

var (
	configFactory ConfigManagerFactory
	configFactoryOnce sync.Once
)

type factoryManagerImpl struct {
	cType ConfigManagerType
}

func ConfigFactory() ConfigManagerFactory {
	configFactoryOnce.Do(func() {
		defaultType := CONFIG_NAIVE
		if cfgType := os.Getenv(util.ENV_CONFIG_IMPL); cfgType != "" {
			defaultType = ConfigManagerType(cfgType)
		}
		configFactory = &factoryManagerImpl{cType: defaultType}
	})
	return configFactory
}

func (f *factoryManagerImpl) SetType(cType ConfigManagerType) ConfigManagerFactory {
	f.cType = cType
	return f
}

func (f *factoryManagerImpl) Build() (ConfigManager, error) {
	switch f.cType {
	case CONFIG_NAIVE, "":
		return &naiveConfigManager{config: make(map[string]interface{})}, nil
	default:
		return nil, fmt.Errorf("ERROR|config/config.Build()|bad type %s for config manager", f.cType)
	}
}