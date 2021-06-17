package util

var (
	Config ConfigMap
)

type ConfigMap struct {
	Host string `toml:"host" mapstructure:"host"`
	Port string `toml:"port" mapstructure:"port"`
	Discord struct {
		Token string `toml:",token" mapstructure:",token"`
		Guild string `toml:",guild" mapstructure:",guild"`
	} `toml:"discord" mapstructure:"discord"`
}