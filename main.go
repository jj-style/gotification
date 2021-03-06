package main

import (
	"bufio"
	"fmt"
	"gotification/src/notify"
	. "gotification/src/util"
	"gotification/src/webservice"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	// set some defaults
	viper.SetDefault("host", "0.0.0.0")
	viper.SetDefault("port", 8080)

	// config from file
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	// config from environment
	viper.AutomaticEnv()
	viper.SetEnvPrefix("ENV")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// config from flag
	pflag.StringP("file", "f", "", "file to read config from")
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		log.Fatal(err)
	}

	// read in the config
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// config file not found
			log.Print(err)
		}
	}
	if viper.GetString("file") != "" {
		loadConfigFromFile()
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatal(err)
	}
}

func loadConfigFromFile() {
	var config string
	if file := viper.GetString("file"); file == "-" {
		lines := make([]string, 0, 100)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		config = strings.Join(lines, "\n")
	} else {
		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		config = string(bytes)
	}
	err := viper.MergeConfig(strings.NewReader(config))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if !Config.Discord.Disable {
		// initialise connection to discord so first request isn't slow
		notify.Discord()
	}

	r := webservice.SetupRouter()
	serverAddr := fmt.Sprintf("%s:%d", Config.Host, Config.Port)
	err := r.Run(serverAddr)
	if err != nil {
		log.Fatal(err)
	}
}
