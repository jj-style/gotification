package main

import (
	"encoding/json"
	"flag"
	"gotify/config"
	"gotify/notify"
	"gotify/webservice"
	"io/ioutil"
	"log"
)

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "config-file", "", "Configuration map")
	flag.Parse()
}

func main() {
	if configFile != "" {
		loadConfig()
	}

	// initialise connection to discord so first request isn't slow
	notify.Discord()

	r := webservice.SetupRouter()
	r.Run()
}

func loadConfig() {
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	var cfg map[string]interface{}
	err = json.Unmarshal(content, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = config.Config().SetMap(cfg)
	if err != nil {
		log.Fatal(err)
	}
}