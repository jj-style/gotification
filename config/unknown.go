package config

import (
	"log"
)

type unknownConfigManager struct {}

func (u *unknownConfigManager) Type() ConfigManagerType {
	return CONFIG_UNKNOWN
}

func (u *unknownConfigManager) SetValue(key string, value interface{}) error {
	log.Fatalf("Unknown config manager has no SetValue")
	return nil
}

func (u *unknownConfigManager) GetString(key string) string {
	log.Fatalf("Unknown config manager has no GetString")
	return ""
}

func (u *unknownConfigManager) SetMap(m map[string]interface{}) error {
	log.Fatalf("Unknown config manager has no SetMap")
	return nil
}