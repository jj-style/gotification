package util

var (
	Config          ConfigMap
	DISABLE_DISCORD = false
)

type ConfigMap struct {
	Host    string           `toml:"host" mapstructure:"host"`
	Port    int              `toml:"port" mapstructure:"port"`
	Discord DiscordConfigMap `toml:"discord" mapstructure:"discord"`
}

type DiscordConfigMap struct {
	Token string `toml:",token" mapstructure:",token"`
	Guild string `toml:",guild" mapstructure:",guild"`
}
