package config

type ConfigManager interface {
	Type() ConfigManagerType
	SetValue(key string, value interface{}) error
	GetString(key string) string
	SetMap(m map[string]interface{}) error
}

type ConfigManagerFactory interface {
	SetType(cType ConfigManagerType) ConfigManagerFactory
	Build() (ConfigManager, error)
}