package util

import "github.com/gin-gonic/gin"

var (
	Config          ConfigMap
	DISABLE_DISCORD = false
)

type ConfigMap struct {
	Host    string        `toml:"host" mapstructure:"host"`
	Port    int           `toml:"port" mapstructure:"port"`
	Auth    AuthConfig    `toml:"auth" mapstructure:"auth"`
	Discord DiscordConfig `toml:"discord" mapstructure:"discord"`
	Extract ExtractConfig `toml:"extract" mapstructure:"extract"`
}

type DiscordConfig struct {
	Token   string `toml:",token" mapstructure:",token"`
	Guild   string `toml:",guild" mapstructure:",guild"`
	Disable bool   `toml:",disable" mapstructure:",disable"`
	NoAuth  bool   `toml:",noauth" mapstructure:",noauth"`
}

type AuthConfig struct {
	Type     string `toml:",type" mapstructure:",type"`
	Accounts []AccountConfig
}

type AccountConfig struct {
	Username string `toml:",username" mapstructure:",username"`
	Password string `toml:",password" mapstructure:",password"`
}

func (c *ConfigMap) GinAccounts() gin.Accounts {
	accounts := make(gin.Accounts, len(c.Auth.Accounts))
	for _, user := range c.Auth.Accounts {
		accounts[user.Username] = user.Password
	}
	return accounts
}

type ExtractConfig struct {
	Type string `toml:",type" mapstructure:",type"`
	Url  string `toml:",url" mapstructure:"url"`
}
